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
	Name  string `bson:"name"`
	Phone int64  `bson:"phone"`
	Email string `bson:"email"`
}
type Activity struct {
	Useractivity string `bson :"useractivity"`
	Duration     string `bson :"duration"`
	AddedTime    string `bson:"addedtime"`
	Email        string `bson:email`
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
	query := bson.M{
		"email": user.Email,
	}
	var data []User
	respon, err := collection.Find(context.Background(), query)
	handleError(err)
	var result string
	respon.All(context.Background(), &data)
	if len(data) != 0 {
		result = "email already present"
	} else {
		_, err = collection.InsertOne(context.Background(), user)
		if err != nil {
			fmt.Println("err", err)
		} else {
			result = "user added"
		}
	}
	w.Write([]byte(result))
}
func Finduser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user = &User{}
	json.NewDecoder(r.Body).Decode(user)
	query := bson.M{
		"email": user.Email,
	}
	// fmt.Println("the query is %v", query)
	err := collection.FindOne(context.Background(), query)
	var result string
	if err != nil {
		fmt.Println("err is ", err)
		result = "couldnot find the user"
	} else {
		result = "user found"
	}

	w.Write([]byte(result))
}
func AddActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var activityofuser = &Activity{}
	json.NewDecoder(r.Body).Decode(activityofuser)
	email := activityofuser.Email
	fmt.Println("email of the user is ", email)
	query := bson.M{
		"email": email,
	}
	cursor, err := collection.Find(context.Background(), query)
	var result string
	handleError(err)
	var result_data []User
	cursor.All(context.Background(), &result_data)
	fmt.Println(result_data)
	if len(result_data) == 0 {
		// ErrNoDocuments means that the filter did not match any documents in
		result = "email doesnt exist"
	} else {
		_, err := collection1.InsertOne(context.Background(), activityofuser)
		if err != nil {
			fmt.Println("error is ", err)
		} else {
			result = "activity added for the user"
		}
	}
	w.Write([]byte(result))
}
func UpdateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var act Activity
	err := json.NewDecoder(r.Body).Decode(&act)
	var result string
	if err != nil {
		fmt.Println("err is %v", err)
		result = "couldnot update the activity"
	} else {
		email := act.Email
		query := bson.M{
			email: email,
		}
		data := bson.M{
			"$set": bson.M{
				"useractivity": act.Useractivity,
				"duration":     act.Duration,
				"addedtime":    act.AddedTime,
			},
		}
		collection1.FindOneAndReplace(context.Background(), query, data)
		result = "activity of the user updated"
	}
	w.Write([]byte(result))
}
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
	myRouter.HandleFunc("/updateactivity", UpdateActivity).Methods("POST")
	myRouter.HandleFunc("/find", Finduser).Methods("GET")
	log.Fatal(http.ListenAndServe("0.0.0.0:10000", myRouter))
}
