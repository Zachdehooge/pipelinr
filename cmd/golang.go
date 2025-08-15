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

var golangCmd = &cobra.Command{
	Use:   "golang",
	Short: "starter pipeline for golang projects",
	Long:  `A starter pipeline with a formatter, builder, linter and security tests for your golang project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/workflows", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		color.Cyan("Downloading golang.yml workflow...")
		if err := DownloadYAMLFromGitHub("https://raw.githubusercontent.com/Zachdehooge/pipelinr/main/cmd/templates/golang.yml", ".github/workflows/golang.yml"); err != nil {
			fmt.Printf("Error downloading golang.yml: %v\n", err)
		}

		color.Cyan("Downloading golang_format.yml workflow...")
		if err := DownloadYAMLFromGitHub("https://raw.githubusercontent.com/Zachdehooge/pipelinr/main/cmd/templates/golang_format.yml", ".github/workflows/golang_format.yml"); err != nil {
			fmt.Printf("Error downloading golang_format.yml: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(golangCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// golangCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// golangCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
