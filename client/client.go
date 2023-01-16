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
	"google.golang.org/protobuf/types/known/timestamppb"
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
	// Adduser(c, "Kanna", "kannavish123@gmail.com", 123456678)
	// AddActivity(c, "kannavish123@gmail.com", "read", 902570235)
	//Adduser(c, "sai", "sai123@gmail.com", 9876543211)
	// AddActivity(c, "kannavish123@gmail.com", "eat",)
	//UpdateActivites(c, "sai123@gmail.com", 5, "sleep")
	//Adduser(c, "hlo", "hlo@gmail.com", 788998668)
	//UpdateActivites(c, "hlo@gmail.com", "sleep")
	//Find(c, "hanshu@gmail.com")
	//Adduser(c, "Hanshu", "hanshu@gmail.com", 123456789)
	//AddActivity(c, "hanshu@gmail.com", "sleep", 9)
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

func Adduser(c pb.TrackerClient, req pb.AddUserRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req1 := &pb.AddUserRequest{User: &pb.User{Name: req.User.Name, Email: req.User.Email, Phone: req.User.Phone}} //Activity: pb.Activity(pb.Activity_value[act.String()])}
	res, err := c.AddUser(ctx, req1)
	if err != nil {
		log.Fatalf("%v errr %v", c, err)
	}
	log.Println(res)
}
func AddActivity(c pb.TrackerClient, email string, activity string, duration int64) {
	reqEmail := email
	reqActivity := activity
	reqduration := duration
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.AddActivityReq{Email: reqEmail, Activity: pb.Activity(pb.Activity_value[reqActivity]), AddedTime: timestamppb.Now(), Duration: reqduration}
	ress, err := c.AddActivity(ctx, req)
	if err != nil {
		log.Fatalf("Err %v", err)
	}
	fmt.Printf("res: %v\n", ress.Response)
}
func UpdateActivites(c pb.TrackerClient, email string, duration int64, activity string) {
	reqEmail := email
	reqActivity := activity
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UpdateActivityReq{Email: reqEmail, Duration: duration, AddedTime: timestamppb.Now(), Activity: reqActivity}
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
