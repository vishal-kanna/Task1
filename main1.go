package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	user := pb.User{}

// 	// we decode our body request params
// 	_ = json.NewDecoder(r.Body).Decode(&user)
// 	// insert our book model.
// 	result, err := collection.InsertOne(context.TODO(), book)

// 	if err != nil {
// 		helper.GetError(err, w)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)
// }
// func handleRequests() {

// }
func handleError(err error) {
	if err != nil {
		log.Fatalf("error is %v", err)
	}
}

type User struct {
	Name    string     `json:"name" bson:"name"`
	Phone   int64      `json:"phone" bson:"phone"`
	Email   string     `json:"email" bson:"email"`
	Act_rec []Activity `json :"act_rec" bson:"act_rec"`
}
type Activity struct {
	Useractivity string `json :"useractivity" bson:"useractivity"`
	Duration     string `json :"duration" bson:"duration"`
	AddedTime    string `json:"added_time" bson:"added_time"`
	Email        string `json:"email" bson:"email"`
}

var collection *mongo.Collection

var collection1 *mongo.Collection

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("I'm inside create user")
	w.Header().Set("Content-Type", "application/json")
	var user = &User{}
	json.NewDecoder(r.Body).Decode(user)
	acts := []Activity{}
	user.Act_rec = acts
	// fmt.Println("user name is", user.Name)
	query := bson.M{
		"email": user.Email,
	}
	var data []User
	respon, err := collection.Find(context.Background(), query)
	handleError(err)
	// var result string
	respon.All(context.Background(), &data)
	if len(data) != 0 {
		// result = "email already present"
		result := "email already present"
		w.Write([]byte(result))

	} else {
		res, err := collection.InsertOne(context.Background(), user)
		if err != nil {
			fmt.Println("err", err)
		}
		json.NewEncoder(w).Encode(res)
	}
	// w.Write([]byte(result))
}
func Finduser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user = &User{}
	json.NewDecoder(r.Body).Decode(user)
	query := bson.M{
		"email": user.Email,
	}
	err := collection.FindOne(context.Background(), query).Decode(user)
	// fmt.Println("user foun", user)
	if err != nil {
		fmt.Println("err is ", err)
	}
	json.NewEncoder(w).Encode(user)
}
func AddActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activityofuser = &Activity{}
	json.NewDecoder(r.Body).Decode(activityofuser)
	fmt.Println("activity", activityofuser)
	email := activityofuser.Email
	fmt.Println("email of the user is ", email)
	query := bson.M{
		"email": email,
	}
	user := &User{}
	collection.FindOne(context.Background(), query).Decode(user)
	fmt.Println("user is ..", user)
	user.Act_rec = append(user.Act_rec, *activityofuser)
	data := bson.M{
		"$set": bson.M{
			"act_rec": user.Act_rec,
		},
	}

	// user1 := User{}
	err := collection.FindOneAndUpdate(context.Background(), query, data).Decode(user)
	if err != nil {
		log.Fatalf("err is %v", err)
	}
	if err == mongo.ErrNoDocuments {
		fmt.Println("err is", err)
	}
	json.NewEncoder(w).Encode(user)
}

// func UpdateActivity(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	// var act Activity
// 	// err := json.NewDecoder(r.Body).Decode(&act)
// 	// var result string
// 	// if err != nil {
// 	// 	fmt.Println("err is %v", err)
// 	// 	result = "couldnot update the activity"
// 	// }
// }
func main() {
	flag.Parse()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	fmt.Println("MongoDB connected")
	err = client.Connect(context.TODO())
	handleError(err)
	collection = (*mongo.Collection)(client.Database("restapi").Collection("Tracker"))
	collection1 = (*mongo.Collection)(client.Database("restapi").Collection("Activity"))
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, os.Interrupt)
	// <-ch
	// fmt.Println("Closing Mongo Connection")
	// if err := client.Disconnect(context.TODO()); err != nil {
	// 	handleError(err)
	// }

	fmt.Println("collection", collection.Name())
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/createuser", CreateUser).Methods("POST")
	myRouter.HandleFunc("/addactivity", AddActivity).Methods("POST")
	// myRouter.HandleFunc("/updateactivity", UpdateActivity).Methods("POST")
	myRouter.HandleFunc("/find", Finduser).Methods("GET")
	log.Fatal(http.ListenAndServe("0.0.0.0:10000", myRouter))
}
