/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// createFolderCmd represents the createFolder command
var createFolderCmd = &cobra.Command{
	Use:   "createFolder [folder name]",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("hello")
		folderName := args[0]
		createFolder(folderName)
	},
}

func createFolder(folderName string) {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	// Create the folder in the working directory
	folderPath := wd + "/" + folderName
	err = os.Mkdir(folderPath, 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	color.Green("Folder '%s' created successfully in %s.\n", folderName, wd)
}

func init() {
	rootCmd.AddCommand(createFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
