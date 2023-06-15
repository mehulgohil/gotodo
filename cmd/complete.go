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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a todo task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := cmd.Flags().GetInt("id")
		if err != nil {
			log.Fatal(err)
		}

		err = service.CompleteTodoTask(taskID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task successfully completed.")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
	completeFlags := completeCmd.Flags()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	completeFlags.Int("id", 0, "ID of the task")
	cobra.MarkFlagRequired(completeFlags, "id")
}
