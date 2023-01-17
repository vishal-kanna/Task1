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

func Main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTrackerClient(conn)
	fmt.Println(c)
}

func Adduser(c pb.TrackerClient, username string, email string, phone int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.AddUserRequest{
		User: &pb.User{
			Name:  username,
			Email: email,
			Phone: phone,
		},
	}
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
	fmt.Println("activity .....",pb.Activity(pb.Activity_value[reqActivity]))
	defer cancel()
	req := &pb.AddActivityReq{Email: reqEmail, Activity: pb.Activity(pb.Activity_value[reqActivity]), AddedTime: time.Now().Format("01-02-2006"), Duration: reqduration}
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
	req := &pb.UpdateActivityReq{Email: reqEmail, Duration: duration, AddedTime: time.Now().Format("01-02-2006 15:04:05 Monday"), Activity: reqActivity}
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
