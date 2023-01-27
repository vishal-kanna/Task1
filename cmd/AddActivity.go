/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	helper "task/client"
	pb "task/protos"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// AddActivityCmd represents the AddActivity command
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
			helper.AddActivity(c, remail, activity, duration)
		}
	},
}

func init() {
	RootCmd.AddCommand(addactivityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddActivityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddActivityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
