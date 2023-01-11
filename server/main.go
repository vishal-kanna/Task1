package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	pb "task/protos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)
var records []*pb.Record

type server struct {
	pb.UnimplementedTrackerServer
}

func (s *server) FindActivity(ctx context.Context, req *pb.SearchActivity) (*pb.Record, error) {
	// fmt.Println("request is", req.Activity)
	query := bson.M{
		"activity_type": req.GetActivity().String(),
	}
	// query := bson.M{
	// 	"activity_type": "sleep",
	// }
	fmt.Println("the query is ", query)
	res := collection.FindOne(ctx, query)
	fmt.Println("result of findOne", res)
	user := &User_record{}
	// handleError(err)
	err := res.Decode(user)
	fmt.Println("user_record is ", user)
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("err: %v\n", err)
	handleError(err)
	result := &pb.Record{
		User: &pb.User{
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
		Activity: pb.Activity(pb.Activity_value[user.Activity_type]),
	}
	// fmt.Println(res1)
	// fmt.Println(reflect.TypeOf(res1))
	return result, nil
}
func (s *server) FindUserByActivity(req *pb.SearchActivity, stream pb.Tracker_FindUserByActivityServer) error {
	fmt.Println("Entered into function")
	query := bson.M{
		"activity_type": req.GetActivity().String(),
	}
	fmt.Printf("query: %v\n", query)
	user := &User_record{}
	users := &pb.Record{}
	cursor, err := collection.Find(context.Background(), query)
	// fmt.Printf("res: %v\n", cursor)
	// fmt.Printf("the type of cursor is :%v", reflect.TypeOf(cursor))
	handleError(err)
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(user)
		handleError(err)
		users = &pb.Record{
			User: &pb.User{
				Name:  user.Name,
				Email: user.Email,
				Phone: user.Phone,
			},
			Activity: pb.Activity(pb.Activity_value[user.Activity_type]),
		}
		stream.Send(users)
	}
	if err := cursor.Err(); err != nil {
		log.Fatalf("Error %v", err)
	}
	return nil
}
func (s *server) Find(ctx context.Context, in *pb.SearchName) (*pb.Record, error) {

	query := bson.M{
		"name": in.Uname,
	}

	user := &User_record{}
	// fmt.Println(query)
	// data:=bson.M{
	// 	"$se"
	// }
	res := collection.FindOne(ctx, query)
	// handleError(err)
	err := res.Decode(user)
	handleError(err)
	result := &pb.Record{
		User: &pb.User{
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
		Activity: pb.Activity(pb.Activity_value[user.Activity_type]),
	}
	// fmt.Println("res1", user)
	// for _, user := range records {
	// 	// fmt.Println(user)
	// 	// fmt.Println(in.GetUname())
	// 	if user.User.Name == in.GetUname() {
	// 		res = user
	// 		break
	// 	}
	// }
	// return res, nil
	// fmt.Println(len(res.RemainingBatchLength()))
	return result, nil
}

type User_record struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	Email         string             `bson:"email"`
	Phone         int64              `bson:"phone"`
	Activity_type string             `bson:"activity_type"`
}
type user_details struct {
	name          string `bson:"name"`
	email         string `bson :"Email"`
	phone         int64  `bson:"phone"`
	activity_type string `bson:"activity_type"`
}

func pushRecordToDB(ctx context.Context, user User_record) string {

	email := user.Email
	query := bson.M{
		"email": email,
	}
	var data []user_details
	respon, err := collection.Find(ctx, query)
	handleError(err)

	respon.All(ctx, &data)
	if len(data) != 0 {
		result := "email already present"
		return result
	}

	collection.InsertOne(ctx, user)
	// if err != nil {
	// 	handleError(err)
	// }
	result := "new user added"
	return result
}

// func (s *server) AllUser(ctx context.Context, in *pb.Empty) (pb.Record, error) {
// 	// log.Printf("Received :%v", in)
// 	// for _, rec := range records {
// 	// 	if err := stream.Send(rec); err != nil {
// 	// 		return err
// 	// 	}
// 	// }
// 	// return nil
// 	query := bson.M{}
// 	cursor, err := collection.Find(context.Background(), query)
// 	handleError(err)
// 	var records []*pb.Record
// 	for cursor.Next(ctx) {
// 		record := &pb.Record{}
// 		err := cursor.Decode(record)
// 		handleError(err)
// 		records = append(records, record)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	if len(records) == 0 {
// 		return records, nil
// 	}

// 	return records, nil

// }
func (s *server) AddUser(ctx context.Context, in *pb.Record) (*pb.Response, error) {
	newUser := User_record{
		Name:          in.User.Name,
		Email:         in.User.Email,
		Phone:         in.User.Phone,
		Activity_type: in.Activity.String(),
	}
	result := pushRecordToDB(ctx, newUser)
	res := &pb.Response{
		Response: result,
	}
	// records = append(records, in)
	return res, nil
}
func (s *server) Update(ctx context.Context, in *pb.RecordReq) (*pb.Response, error) {
	// res := pb.Record{}
	res := pb.Response{}
	// user_name := in.User.Name
	user_email := in.User.Email
	// user_phone := in.User.Phone
	user_Activity_type := in.GetActivity()
	// collection.ReplaceOne(ctx ,{"Name":user_name },{"Activity_type":useruser_Activity_type});
	// updated_data := User_record{
	// 	Name:          user_name,
	// 	Email:         user_email,
	// 	Phone:         user_phone,
	// 	Activity_type: user_Activity_type.String(),
	// }
	// d := bson.D{}
	data := bson.M{
		"$set": bson.M{
			// "name":          user_name,
			// "email":         user_email,
			// "phone":         user_phone,
			"activity_type": user_Activity_type.String(),
		},
	}
	query := bson.M{
		"email": user_email,
	}

	// fmt.Println("query", query)
	// fmt.Println("data", data)
	_, err := collection.UpdateOne(ctx, query, data)
	handleError(err)
	// fmt.Println("the response is ", resp.ModifiedCount, err)
	return &res, nil

	// res = pb.Record{User: &pb.User{Name: user_name, Email: user_email, Phone: user_phone}, Activity: pb.Activity(pb.Activity_value[user_Activity_type])}
	// return &res, nil
	// res := &pb.Record{}
	// user_name := in.User.Name
	// fmt.Println("the user whose activity is to be updated is ", in) // activity_update :=in.Activity

	// for i := 0; i < len(records); i++ {
	// 	if records[i].User.Name == user_name {
	// 		fmt.Println("the record before updating the activity", records[i])
	// 		newREc := &pb.Record{User: &pb.User{Name: in.User.Name, Email: records[i].User.Email, Phone: records[i].User.Phone}, Activity: pb.Activity(pb.Activity_value[in.GetActivity()])}
	// 		// records[i].Activity = //pb.Activity(pb.Activity_value[in.GetActivity()])
	// 		// records[i].Activity = pb.Activity_play
	// 		// fmt.Println("updating the activity ", records[i].Activity)
	// 		// fmt.Println("the record after updating the activity1", records[i].Activity)

	// 		// fmt.Println("the record after updating the activity", records[i])
	// 		// res = records[i]
	// 		pushRecordToDB(ctx, records[i])
	// 		records[i] = newREc
	// 		// records[i]=
	// 		res = newREc
	// 		break
	// 	}
	// }
	// // for _, record := range records {
	// // 	if record.User.Name == user_name {
	// // 		fmt.Println("matching record", record)
	// // 		fmt.Printf("in.GetActivity(): %v\n", in.GetActivity())
	// // 		fmt.Printf("pb.Activity(pb.Activity_value[in.GetActivity()]): %v\n", pb.Activity(pb.Activity_value[in.GetActivity()]))
	// // 		record.Activity = pb.Activity(pb.Activity_value[in.GetActivity()])
	// // 		newRec := &pb.Record{
	// // 			User:&pb.User{Name:in.User.Name,Email: in.User.Email,}
	// // 			Activity: pb.Activity(pb.Activity_value[in.GetActivity()]),
	// // 		}

	// // 		fmt.Println("newRec", newRec)
	// // 		fmt.Println("updated record", record)
	// // 		res = record
	// // 		break
	// // 	}
	// // }
	// fmt.Println("The updated activity of the user is ", res.GetActivity())
	// fmt.Println("resulting activity", res)
	// return res, nil
}

// func initRecords() {
// 	// movie1 := &pb.MovieInfo{Id: "1", Name: "Vishal", Isbn: "0987655432", Rating: "4"}
// 	// movie1 := &pb.MovieInfo{Id: "1", Isbn: "0987655432", Title: "The Batman", Director: &pb.Director{Fname: "Sam", Lname: "Son"}}

// 	record1 := &pb.Record{User: &pb.User{Name: "Vishal", Email: "kannavish", Phone: 7095896613}, Activity: pb.Activity_eat}
// 	record2 := &pb.Record{User: &pb.User{Name: "Kanna", Email: "vitwit", Phone: 6301244842}, Activity: pb.Activity_sleep}
// 	records = append(records, record1)
// 	records = append(records, record2)

// }
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var collection *mongo.Collection

func main() {
	flag.Parse()
	// initRecords()
	s := grpc.NewServer()
	pb.RegisterTrackerServer(s, &server{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server Listening at %v", lis.Addr())
	go func() {

		if err := s.Serve(lis); err != nil {
			// log.Fatalf("Failed to serve :%v", err)
			handleError(err)
		}
	}()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	fmt.Println("MongoDB connected")
	err = client.Connect(context.TODO())
	handleError(err)
	collection = (*mongo.Collection)(client.Database("Record").Collection("Record"))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Closing Mongo Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)
	}
	s.Stop()
}
