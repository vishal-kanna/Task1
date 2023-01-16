/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateactivityCmd represents the updateactivity command
var updateactivityCmd = &cobra.Command{
	Use:   "updateactivity",
	Short: "adds the activity for the user",
	Long:  `Updates the activity of the user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateactivity called")
	},
}

func init() {
	rootCmd.AddCommand(updateactivityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateactivityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateactivityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
