/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/mehulgohil/gotodo/service"
	"log"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a todo task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		keyword, err := cmd.Flags().GetString("keyword")
		if err != nil {
			log.Fatal(err)
		}
		duedate, err := cmd.Flags().GetString("duedate")
		if err != nil {
			log.Fatal(err)
		}

		tasks, err := service.SearchTodoTask(keyword, duedate)
		if err != nil {
			return
		}

		if len(tasks.Todos) == 0 {
			fmt.Println("There are no tasks with given search filter.")
			return
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
# Example 1: Search for tasks with keyword and duedate
gotodo search -k te -d 12/11/23

# Example 2: Search for tasks with keyword test
gotodo search -k test

# Example 3: Search for tasks with a duedate
gotodo search -d 12/11/23
	`,
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().StringP("keyword", "k", "", "Keyword for the task name to search for.")
	searchCmd.Flags().StringP("duedate", "d", "", "Duedate of the task to search for.")
}
