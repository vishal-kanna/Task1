package main

import (
	"context"
	"fmt"
	"log"
	pb "task/protos"
	"testing"

	"google.golang.org/grpc"
)

func TestAddUser(t *testing.T) {
	connection, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("the error is ", err)
	}
	c := pb.NewTrackerClient(connection)
	type expectation struct {
		Expresult string
	}
	tests := map[string]struct {
		in       *pb.AddUserRequest
		expected expectation
	}{
		"Most_Success": {
			in: &pb.AddUserRequest{
				User: &pb.User{
					Name:  "vishalkanna",
					Phone: 743599899,
					Email: "kanna1@gmail.com",
				},
			},
			expected: expectation{
				Expresult: "new user added",
			},
		},
		"Already Present": {
			in: &pb.AddUserRequest{
				User: &pb.User{
					Name:  "vishalkanna",
					Phone: 743599899,
					Email: "kanna@gmail.com",
				},
			},
			expected: expectation{
				Expresult: "new user added",
			},
		},
	}
	// got := Adduser(c, "testuser1", "testuser1@gmail.com", 101010)
	// want := "new user added"

	// if got != want {
	// 	t.Errorf("got %q, wanted %q", got, want)
	// }
	for senario, tt := range tests {
		// fmt.Println("hello im in TestAdduser")
		t.Run(senario, func(t *testing.T) {
			fmt.Println("hello im in TestAdduser")
			fmt.Println("the tt.in value is", tt.in.User)
			out, err := c.AddUser(context.Background(), tt.in)
			fmt.Println("heii")
			// fmt.Println("err")
			if err != nil {
				// fmt.Println("hello im in TestAdduser")
				t.Errorf("Err%q", err)
			} else {
				if tt.expected.Expresult == out.Response {
					fmt.Println("tt.expected.Expresult is", tt.expected.Expresult, out.Response)
				} else {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.Expresult, out.Response)
				}
			}

		})
	}
}

func TestFindUser(t *testing.T) {
	connection, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("the error is ", err)
	}
	c := pb.NewTrackerClient(connection)
	got, err := Find(c, "kannavitwit@gmail.com")
	want := pb.FindUserRes{
		User: &pb.User{
			Name:  "kanna",
			Email: "kannavitwit@gmail.com",
			Phone: 743599899,
		},
	}
	if got.User.Name == want.User.Name && got.User.Email == want.User.Email && got.User.Phone == want.User.Phone {
		fmt.Println("Results are same")
	} else {
		t.Errorf("Err%q", err)
	}
}

func TestAddactivity(t *testing.T) {
	connection, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("the error is ", err)
	}
	c := pb.NewTrackerClient(connection)
	res, err := AddActivity(c, "kannavitwit@gmail.com", "read", 40)
	if err != nil {
		log.Fatalf("err is %v", err)
	}
	fmt.Println("response is ", res)
	want := &pb.AddActivityRes{
		Response: "Updated",
	}
	if want.Response == res {
		fmt.Printf("want is %v and got is %v", res, want.Response)
	} else {
		t.Errorf("Err %q", err)
	}
}
func TestUpdateActivity(t *testing.T) {
	connection, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("the error is ", err)
	}
	c := pb.NewTrackerClient(connection)
	res, err := UpdateActivites(c, "kannavitwit@gmail.com", 34, "sleep")
	want := pb.UpdateActivityRes{
		Response: "Updated",
	}
	if res == want.Response {
		fmt.Printf("want is %v and got is %v", res, want.Response)
	} else {
		t.Errorf("Err %q", err)
	}
}
