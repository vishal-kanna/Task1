/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	helper "task/client"
	pb "task/protos"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// FindUserCmd represents the FindUser command
var FinduserCmd = &cobra.Command{
	Use:   "Finduser",
	Short: "Finds the user",
	Long:  `Finds the user of the given email`,
	Run: func(cmd *cobra.Command, args []string) {
		connection, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println("the error is ", err)
		}
		c := pb.NewTrackerClient(connection)
		if len(args) <= 0 {
			fmt.Println("Nothing is passed in the cmd")
		} else {
			helper.Find(c, args[0])
			// helper.Find(c, args[0])
		}
	},
}

func init() {
	RootCmd.AddCommand(FinduserCmd)
}
