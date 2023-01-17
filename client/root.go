/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"os"
	"strconv"
	pb "task/protos"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

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

		Adduser(c, username, email, phone)
	},
}

var addactivityCmd = &cobra.Command{
	Use:   "addactivity",
	Short: "adds the activity to the user",
	Long: `Adds the activity to the user and also adds the duration of the activity
	add activity takes the email and finds the user and adds the activity,duration to the user
	
	
		email activity duration
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("addactivity called")
		connection, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println("the error is ", err)
		}
		c := pb.NewTrackerClient(connection)

		remail := args[0]
		activity := args[1]
		duration, _ := strconv.ParseInt(args[2], 10, 64)
		// fmt.Println("email ...activity ....duration", remail, activity, duration)
		if len(args) <= 2 {
			fmt.Println("the arguments are not sufficient ........email ,activity,duration")
		} else {
			AddActivity(c, remail, activity, duration)
		}
	},
}
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
			Find(c, args[0])
		}
	},
}

var updateactivityCmd = &cobra.Command{
	Use:   "updateactivity",
	Short: "adds the activity for the user",
	Long:  `Updates the activity of the user`,
	Run: func(cmd *cobra.Command, args []string) {
		connection, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			fmt.Println("the error is ", err)
		}
		c := pb.NewTrackerClient(connection)
		if len(args) <= 0 {
			fmt.Println("pls enter the email,activity and duration of the activity of the user to update  the activity")
		} else {
			remail := args[0]
			activity := args[1]
			duration, _ := strconv.ParseInt(args[2], 10, 64)
			UpdateActivites(c, remail, duration, activity)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(createuserCmd)
	rootCmd.AddCommand(addactivityCmd)
	rootCmd.AddCommand(FinduserCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
