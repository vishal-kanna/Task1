/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	pb "task/protos"

	helper "task/client"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// CreateUserCmd represents the CreateUser command
var createuserCmd = &cobra.Command{
	Use:   "createuser",
	Short: "Add user ",
	Long:  `Adding a new user on the server using Grpc`,
	Run: func(cmd *cobra.Command, args []string) {
		connection, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println("the error is ", err)
		}
		c := pb.NewTrackerClient(connection)
		// fmt.Println("createuser called")
		fmt.Println(args)
		username := args[0]

		email := args[1]

		phone, _ := strconv.ParseInt(args[2], 10, 64)

		helper.Adduser(c, username, email, phone)
	},
}

func init() {
	RootCmd.AddCommand(createuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CreateUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CreateUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
