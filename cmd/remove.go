/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mehulgohil/gotodo/service"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo tasks",
	Args:  cobra.ExactArgs(1), //accepts 1 args
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}

		err = service.RemoveTodoTask(taskID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task successfully removed")
	},
	Example: `
# Example 1: Remove a task with id 1
gotodo remove 1
	`,
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
