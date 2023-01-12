package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	pb "task/protos"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	// name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTrackerClient(conn)
	// runFindUser(c, "Kanna")
	// runGetAllUsers(c)
	//Adduser(c, "Kanna", "kanna123@gmail.com", 123456789, pb.Activity_sleep)
	//Adduser(c, "Vishal", "vitwit@gmail.com", 2233456789, pb.Activity_read)
	//Update(c, "vitwit@gmail.com", pb.Activity_play)
	// runGetAllUsers(c)

	//runFindUser(c, "Kanna")
	//runAdduser(c, "Gunni", "aakash@tsbpass", 123456789, pb.Activity_eat)
	// runGetAllUsers(c)
	// runFindUser(c, "Vishal")
	// GetAllUsers(c)
	// Update(c, "Vishal", pb.Activity_read)
	// runFindUser(c, "Vishal")
	// Update(c, "aakash@tsbpass", pb.Activity_play)
	// GetAllUsers(c)
	// FindUser(c, "Kanna")
	//Adduser(c, "SaiTeja", "saiteja@vitwit", 123456789, pb.Activity_read)
	// GetActivityUser(c, "sleep")
	// GetAllUsersByActivity(c, "read")
	//Adduser(c, "Vishal", "Vishal123@gmail.com", 123456789)
	//AddActivity(c, "Vishal123@gmail.com", "eat")
	//Adduser(c, "Kanna", "kannavish123@gmail.com", 123456678)
	//AddActivity(c, "kannavish123@gmail.com", "read")
	//Adduser(c, "sai", "sai123@gmail.com", 9876543211)
	AddActivity(c, "kannavish123@gmail.com", "eat")
	//UpdateActivites(c, "sai123@gmail.com", "sleep")
	//Adduser(c, "hlo", "hlo@gmail.com", 788998668)
	//UpdateActivites(c, "hlo@gmail.com", "sleep")
	// Find(c, "hlo@gmail.com")
}

// func (s* server)
// func FindUser(c pb.TrackerClient, name string) {
// 	flag.Parse()
// 	// fmt.Println(name)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.SearchName{Uname: name}
// 	res, err := c.Find(ctx, req)
// 	if err != nil {
// 		log.Fatalf("Error %v", err)
// 	}
// 	log.Printf("The Queried user is  %v", res)
// }

func Adduser(c pb.TrackerClient, name string, email string, phonenum int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.AddUserRequest{User: &pb.User{Name: name, Email: email, Phone: phonenum}} //Activity: pb.Activity(pb.Activity_value[act.String()])}
	res, err := c.AddUser(ctx, req)
	if err != nil {
		log.Fatalf("%v errr %v", c, err)
	}
	log.Println(res)
}
func AddActivity(c pb.TrackerClient, email string, activity string) {
	reqEmail := email
	reqActivity := activity
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.AddActivityReq{Email: reqEmail, Activity: pb.Activity(pb.Activity_value[reqActivity])}
	_, err := c.AddActivity(ctx, req)
	if err != nil {
		log.Fatalf("Err %v", err)
	}
	// fmt.Printf("res: %v\n", res.Response)
}
func UpdateActivites(c pb.TrackerClient, email string, activity string) {
	reqEmail := email
	reqActivity := activity
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UpdateActivityReq{Email: reqEmail, Activity: reqActivity}
	res, err := c.UpdateActivites(ctx, req)
	if err != nil {
		log.Fatalf("Err %v", err)
	}
	fmt.Println("Result is ", res)
}
func Find(c pb.TrackerClient, email string) {
	reqEmail := email
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.FindUserReq{Email: reqEmail}
	res, err := c.Find(ctx, req)
	if err != nil {
		fmt.Println("error is ", err)
	}
	fmt.Println("result is ", res)
}

// // func Update(c pb.TrackerClient, email string, act pb.Activity) {
// // 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// // 	defer cancel()
// // 	req := &pb.RecordReq{User: &pb.User{Email: email}, Activity: pb.Activity(pb.Activity_value[act.String()])}
// // 	_, err := c.Update(ctx, req)
// // 	if err != nil {
// // 		log.Fatalf("Error in Updatin the user %v", err)
// // 	}
// // }

// func GetActivityUser(c pb.TrackerClient, activity string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	act := &pb.SearchActivity{Activity: pb.Activity(pb.Activity_value[activity])}
// 	res, err := c.FindActivity(ctx, act)
// 	if err != nil {
// 		log.Fatal("Error occured", err)
// 	}
// 	fmt.Println(res)
// }
// func GetAllUsersByActivity(c pb.TrackerClient, activity string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	req := &pb.SearchActivity{Activity: pb.Activity(pb.Activity_value[activity])}
// 	stream, err := c.FindUserByActivity(ctx, req)
// 	if err != nil {
// 		log.Printf("error %v", err)
// 	}
// 	for {
// 		// stream.Recv returns a pointer to a Records at the current iteration
// 		res, err := stream.Recv()
// 		// If end of stream, break the loop
// 		if err == io.EOF {
// 			break
// 		}
// 		// if err, return an error
// 		if err != nil {
// 			return err
// 		}
// 		// If everything went well use the generated getter to print the Record message
// 		fmt.Println(res)
// 	}
// 	return err
// }
