/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/mehulgohil/gotodo/models"
	"github.com/mehulgohil/gotodo/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var jsonFile string

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load the existing todo task json",
	Long:  ``,
	// START OMIT
	Run: func(cmd *cobra.Command, args []string) {
		var existingTasks models.Todos

		viper.SetConfigFile(jsonFile) // HL
		err := viper.ReadInConfig()   // HL
		if err != nil {               // HL
			log.Fatal(err) // HL
		} // HL

		err = viper.Unmarshal(&existingTasks) // HL
		if err != nil {                       // HL
			log.Fatal(err) // HL
		} // HL

		err = service.LoadExistingJSON(existingTasks)
		if err != nil {
			log.Fatal(err)
		}
	},
	// END OMIT
	Example: `
# Example 1: Remove a task with id 1
gotodo load --jsonFile existingFile.json
	`,
}

func init() {
	rootCmd.AddCommand(loadCmd)
	loadFlags := loadCmd.Flags()

	loadFlags.StringVar(&jsonFile, "jsonFile", "", "json file for existing tasks"+
		"")
	cobra.MarkFlagRequired(loadFlags, "jsonFile")
}
