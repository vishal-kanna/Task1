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

// FinduserCmd represents the Finduser command
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
		} // fmt.Println("Finduser called")
		//fmt.Println("the args [] array is ", args[1])
	},
}

func init() {

	rootCmd.AddCommand(FinduserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// FinduserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// FinduserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
