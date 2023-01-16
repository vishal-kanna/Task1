/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	pb "task/protos"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// createuserCmd represents the createuser command
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
		username, _ := cmd.Flags().GetString("username")

		email, _ := cmd.Flags().GetString("email")

		phone, _ := cmd.Flags().GetInt64("phone")
		// if err != nil {
		// 	return err
		// }
		//user proto
		user := &pb.User{
			Name:  username,
			Email: email,
			Phone: phone,
		}
		// name, _ := cmd.Flags().GetString("username")
		// email, _ := cmd.Flags().GetString("email")

		// Rpc call
		req := &pb.AddUserRequest{
			User: &pb.User{
				Name:  username,
				Email: email,
				Phone: phone,
			},
		}
		res, err := Adduser(c, req)
		if err != nil {
			return err
		}
		fmt.Println("the user is added", user)
	},
}

func init() {
	// rootCmd.AddCommand(createuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createuserCmd.Flags().StringP("username", "u", "", "username of the user")
	createuserCmd.Flags().StringP("email", "e", "", "Email of the user")
	createuserCmd.Flags().Int64P("phone", "p", 0, "The Phone number of the user")
	createuserCmd.MarkFlagRequired("username")
	createuserCmd.MarkFlagRequired("email")
	createuserCmd.MarkFlagRequired("phone")
	rootCmd.AddCommand(createuserCmd)
}
