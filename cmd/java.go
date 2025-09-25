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

// javaCmd represents the java command
var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "starter pipeline for java",
	Long:  `A starter pipeline with a formatter, builder, linter and security tests for your java project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/workflows", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		color.Red("Downloading java.yml workflow...")
		if err := DownloadYAMLFromGitHub("https://raw.githubusercontent.com/Zachdehooge/pipelinr/main/cmd/templates/java.yml", ".github/workflows/java.yml"); err != nil {
			fmt.Printf("Error downloading java.yml: %v\n", err)
		}

		color.Red("Downloading java_format.yml workflow...")
		if err := DownloadYAMLFromGitHub("https://raw.githubusercontent.com/Zachdehooge/pipelinr/main/cmd/templates/java_format.yml", ".github/workflows/java_format.yml"); err != nil {
			fmt.Printf("Error downloading java_format.yml: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(javaCmd)
}
