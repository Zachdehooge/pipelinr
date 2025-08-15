/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "starter pipeline for python",
	Long:  `A starter pipeline with a formatter, builder, linter and security tests for your python project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/workflows", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		color.Yellow("Downloading python.yml workflow...")
		if err := DownloadYAMLFromGitHub("https://raw.githubusercontent.com/Zachdehooge/pipelinr/main/cmd/templates/python.yml", ".github/workflows/python.yml"); err != nil {
			fmt.Printf("Error downloading python.yml: %v\n", err)
		}

		color.Yellow("Downloading python_format.yml workflow...")
		if err := DownloadYAMLFromGitHub("https://raw.githubusercontent.com/Zachdehooge/pipelinr/main/cmd/templates/python_format.yml", ".github/workflows/python_format.yml"); err != nil {
			fmt.Printf("Error downloading python_format.yml: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pythonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pythonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
