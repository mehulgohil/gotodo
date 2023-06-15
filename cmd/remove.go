/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mehulgohil/gotodo/service"
	"log"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := cmd.Flags().GetInt("id")
		if err != nil {
			log.Fatal(err)
		}

		err = service.RemoveTodoTask(taskID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task successfully removed")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	removeCmd.Flags().Int("id", 0, "ID of the task")
}
