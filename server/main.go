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

// var records []*pb.Recor

type server struct {
	pb.UnimplementedTrackerServer
}

// func (s *server) FindActivity(ctx context.Context, req *pb.SearchActivity) (*pb.Record, error) {
// 	query := bson.M{
// 		"activity_type": req.GetActivity().String(),
// 	}
// 	fmt.Println("the query is ", query)
// 	res := collection.FindOne(ctx, query)
// 	fmt.Println("result of findOne", res)
// 	user := &User_record{}
// 	err := res.Decode(user)
// 	fmt.Println("user_record is ", user)
// 	handleError(err)
// 	result := &pb.Record{
// 		User: &pb.User{
// 			Name:  user.Name,
// 			Email: user.Email,
// 			Phone: user.Phone,
// 		},
// 		Activity: pb.Activity(pb.Activity_value[user.Activity_type]),
// 	}
// 	// fmt.Println(res1)
// 	// fmt.Println(reflect.TypeOf(res1))
// 	return result, nil
// }
// func (s *server) FindUserByActivity(req *pb.SearchActivity, stream pb.Tracker_FindUserByActivityServer) error {
// 	fmt.Println("Entered into function")
// 	query := bson.M{
// 		"activity_type": req.GetActivity().String(),
// 	}
// 	fmt.Printf("query: %v\n", query)
// 	user := &User_record{}
// 	users := &pb.Record{}
// 	cursor, err := collection.Find(context.Background(), query)
// 	handleError(err)
// 	defer cursor.Close(context.Background())
// 	for cursor.Next(context.Background()) {
// 		err := cursor.Decode(user)
// 		handleError(err)
// 		users = &pb.Record{
// 			User: &pb.User{
// 				Name:  user.Name,
// 				Email: user.Email,
// 				Phone: user.Phone,
// 			},
// 			Activity: pb.Activity(pb.Activity_value[user.Activity_type]),
// 		}
// 		stream.Send(users)
// 	}
// 	if err := cursor.Err(); err != nil {
// 		log.Fatalf("Error %v", err)
// 	}
// 	return nil
// }
// func (s *server) Find(ctx context.Context, in *pb.SearchName) (*pb.Record, error) {

// 	query := bson.M{
// 		"name": in.Uname,
// 	}

// 	user := &User_record{}
// 	res := collection.FindOne(ctx, query)
// 	err := res.Decode(user)
// 	handleError(err)
// 	result := &pb.Record{
// 		User: &pb.User{
// 			Name:  user.Name,
// 			Email: user.Email,
// 			Phone: user.Phone,
// 		},
// 		Activity: pb.Activity(pb.Activity_value[user.Activity_type]),
// 	}
// 	return result, nil
// }

//db model
type User_record struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	Email         string             `bson:"email"`
	Phone         int64              `bson:"phone"`
	Activity_type []string           `bson:"activity_type"`
}

type user_details struct {
	name          string `bson:"name"`
	email         string `bson :"Email"`
	phone         int64  `bson:"phone"`
	activity_type string `bson:"activity_type"`
}

//pushing the data into the database
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
	result := "new user added"
	return result
}
func (s *server) AddUser(ctx context.Context, in *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	newUser := User_record{
		Name:          in.User.Name,
		Email:         in.User.Email,
		Phone:         in.User.Phone,
		Activity_type: []string{},
		// Activity_type: in.Activity.String(),
	}
	result := pushRecordToDB(ctx, newUser)
	res := &pb.AddUserResponse{
		Response: result,
	}
	return res, nil
}

//Adding the Activity to the Particular user
func (s *server) AddActivity(ctx context.Context, req *pb.AddActivityReq) (*pb.AddActivityRes, error) {
	query := bson.M{
		"email": req.Email,
	}
	activity := req.GetActivity().String()
	res1 := collection.FindOne(ctx, query)
	user := &User_record{}
	res1.Decode(user)
	user.Activity_type = append(user.Activity_type, activity)
	data := bson.M{
		"$set": bson.M{
			"activity_type": user.Activity_type,
		},
	}
	result := &pb.AddActivityRes{}
	res := collection.FindOneAndUpdate(ctx, query, data)
	if res != nil {
		log.Fatalf("err is %v", res)
	} else {
		result = &pb.AddActivityRes{
			Response: "Updated",
		}
	}
	return result, nil
}

//Updating the activites of the user
func (s *server) UpdateActivites(ctx context.Context, req *pb.UpdateActivityReq) (*pb.UpdateActivityRes, error) {
	reqEmail := req.Email
	reqActivity := req.Activity
	query := bson.M{
		"email": reqEmail,
	}
	user := &User_record{}
	res := collection.FindOne(ctx, query)
	// fmt.Println("res..", res)
	err := res.Decode(user)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	user.Activity_type = append(user.Activity_type, reqActivity)
	data := bson.M{
		"$set": bson.M{
			"activity_type": user.Activity_type,
		},
	}
	collection.FindOneAndUpdate(ctx, query, data)
	result := &pb.UpdateActivityRes{}
	result = &pb.UpdateActivityRes{
		Response: "Updated",
	}
	return result, nil
}
func (s *server) Find(ctx context.Context, req *pb.FindUserReq) (*pb.FindUserRes, error) {
	reqemail := req.Email
	result := &pb.FindUserRes{}
	query := bson.M{
		"email": reqemail,
	}
	user := &User_record{}
	res := collection.FindOne(ctx, query)
	err := res.Decode(user)
	if err != nil {
		fmt.Println("error is ", err)
	}
	result = &pb.FindUserRes{User: &pb.User{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
	},

	// Activities: []pb.Activity(user.Activity_type[]),
	}
	return result, nil
	//	query := bson.M{
	// 		"name": in.Uname,
	// 	}

	// 	user := &User_record{}
	// 	res := collection.FindOne(ctx, query)
	// 	err := res.Decode(user)
	// 	handleError(err)
	// 	result := &pb.Record{
	// 		User: &pb.User{
	// 			Name:  user.Name,
	// 			Email: user.Email,
	// 			Phone: user.Phone,
	// 		},
	// 		Activity: pb.Activity(pb.Activity_value[user.Activity_type]),
	// 	}
	// 	return result, nil

}

// func (s *server) Update(ctx context.Context, in *pb.RecordReq) (*pb.Response, error) {
// 	res := pb.Response{}

// 	user_email := in.User.Email
// 	user_Activity_type := in.GetActivity()
// 	data := bson.M{
// 		"$set": bson.M{
// 			"activity_type": user_Activity_type.String(),
// 		},
// 	}
// 	query := bson.M{
// 		"email": user_email,
// 	}
// 	_, err := collection.UpdateOne(ctx, query, data)
// 	handleError(err)
// 	return &res, nil
// }
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var collection *mongo.Collection

func main() {
	flag.Parse()
	s := grpc.NewServer()
	pb.RegisterTrackerServer(s, &server{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server Listening at %v", lis.Addr())
	go func() {

		if err := s.Serve(lis); err != nil {
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
