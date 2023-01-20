package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "task/protos"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func Server(ctx context.Context) (pb.TrackerClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterTrackerServer(baseServer, &server{})
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewTrackerClient(conn)

	return client, closer
}
func TestAdduser(t *testing.T) {
	// fmt.Println("hello im in TestAdduser")
	ctx := context.Background()
	client, closer := Server(ctx)
	defer closer()
	// fmt.Println("hello im in TestAdduser")
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
					Name:  "kanna",
					Phone: 743599899,
					Email: "kannavitwit@gmail.com",
				},
			},
			expected: expectation{
				Expresult: "new user added",
			},
		},
	}
	// fmt.Println("hello im in TestAdduser", tests)
	// fmt.Println("hello im in TestAdduser")
	for _, tt := range tests {
		// fmt.Println("hello im in TestAdduser")
		//t.Run(senario, func(t *testing.T) {
		// fmt.Println("hello im in TestAdduser")
		// fmt.Println("the tt.in value is", tt.in.User)
		out, err := client.AddUser(ctx, tt.in)
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

		// })
	}
}
