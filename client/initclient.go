package main

import (
	"context"
	"fmt"
	"log"
	pb "task/protos"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var cfgFile string

// Client and context global vars for the cmd package
// So they can be used by our subcommands
var client pb.TrackerClient
var requestCtx context.Context
var requestOpts grpc.DialOption

func init() {
	// initConfig reads in config file and ENV variables
	cobra.OnInitialize(initConfig)
	// After Cobra root config init, initialize the client
	fmt.Println("Starting User Tracker Server Client")
	// Establish context to timeout after 10 seconds if server does not respond
	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	// Establish insecure grpc options (no TLS)
	requestOpts = grpc.WithInsecure()
	// Dial the server, returns a client connection
	conn, err := grpc.Dial("localhost:50051", requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish client connection to localhost:50051: %v", err)
	}
	// Instantiate the BlogServiceClient with our client connection to the server
	client = pb.NewTrackerClient(conn)
}
