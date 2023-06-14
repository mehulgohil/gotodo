/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mehulgohil/gotodo/models"
	"github.com/mehulgohil/gotodo/service"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskName, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		taskDueDate, err := cmd.Flags().GetString("duedate")
		if err != nil {
			log.Fatal(err)
		}

		currentTodo := models.Todo{
			Name:      taskName,
			DueDate:   taskDueDate,
			Completed: false,
		}

		err = service.AddTodo(currentTodo)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().StringP("name", "n", "", "Name of the task")
	addCmd.Flags().StringP("duedate", "d", "", "Due date of the task")
}
