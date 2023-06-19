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

// completeCmd represents the complete command
// START OMIT
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a todo task",
	Long:  ``,
	Args:  cobra.ExactArgs(1), // HL
	Run: func(cmd *cobra.Command, args []string) {
		// Reading Args 0
		taskID, err := strconv.Atoi(args[0]) // HL
		if err != nil {                      // HL
			log.Fatal(err) // HL
		} // HL

		err = service.CompleteTodoTask(taskID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Task successfully completed.")
	},
	// END OMIT
	Example: `
# Example 1: complete a task
gotodo complete 1
	`,
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
