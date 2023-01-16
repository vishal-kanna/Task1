package main

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

import (
	"fmt"
	pb "task/protos"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// addactivityCmd represents the addactivity command
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

		remail, _ := cmd.Flags().GetString("email")
		activity, _ := cmd.Flags().GetString("activity")
		duration, _ := cmd.Flags().GetInt64("duration")
		// fmt.Println("email ...activity ....duration", remail, activity, duration)
		if len(args) <= 2 {
			fmt.Println("the arguments are not sufficient")
		} else {
			// email :=
			// duration := strconv.Itoa(args[])
			// activity := args[2]
			req := pb.AddActivityReq{
				Email:    remail,
				Duration: duration,
				Activity: pb.Activity(pb.Activity_value[activity]),
			}
			res, err := AddActivity(c, req)
			if err != nil {
				fmt.Println("error is ", err)
			}
		}
	},
}

func init() {
	// addactivityCmd.Flags().StringP("username", "u", "", "username of the user")
	addactivityCmd.Flags().StringP("email", "e", "", "Email of the user")
	addactivityCmd.Flags().Int64P("duration", "d", 7, "duration")
	addactivityCmd.Flags().StringP("activity", "a", "", "activity of the user")
	addactivityCmd.MarkFlagRequired("email")
	addactivityCmd.MarkFlagRequired("duration")
	addactivityCmd.MarkFlagRequired("activity")
	// addactivityCmd.AddCommand(createuserCmd)
	addactivityCmd.AddCommand(addactivityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addactivityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addactivityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
