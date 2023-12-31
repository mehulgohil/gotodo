/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mehulgohil/gotodo/service"
	"github.com/mehulgohil/gotodo/utility"
	"github.com/spf13/cobra"
	"log"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "A subcommand for search to search for all completed tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// using a persistent flag defined in parent command
		outputFormat, err := searchCmd.PersistentFlags().GetString("output")
		if err != nil {
			log.Fatal(err)
		}

		tasks, err := service.SearchCompletedTask()
		if err != nil {
			log.Fatal(err)
		}

		if len(tasks.Todos) == 0 {
			fmt.Println("There are no completed task/s.")
			return
		}

		utility.PrintOnOutputFormat(tasks, outputFormat)
	},
	Example: `
# Example 1: Search for a completed task
gotodo search completed
	`,
}

func init() {
	searchCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
