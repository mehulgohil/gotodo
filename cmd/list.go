/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mehulgohil/gotodo/service"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todo tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := service.GetAllTodoTasks()
		if err != nil {
			log.Fatal(err)
		}
		if len(tasks.Todos) == 0 {
			fmt.Println("There are no task added to todo list.")
		}
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 5, ' ', 0)

		fs := "%s\t%s\t%s\t%s\n"
		fmt.Fprintf(w, fs, "ID", "Name", "Due Date", "Completed")
		for _, eachTask := range tasks.Todos {
			fmt.Fprintf(w, fs, strconv.Itoa(eachTask.ID), eachTask.Name, eachTask.DueDate, strconv.FormatBool(eachTask.Completed))
		}

		w.Flush()
	},
	Example: `
# Example 1: List all tasks
gotodo list
	`,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
