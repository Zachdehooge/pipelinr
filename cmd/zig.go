/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func zig() []byte {
	var text = `name: CI/CD

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest
    name: Build
    steps:
      - uses: actions/checkout@v3
      - uses: mlugg/setup-zig@v2
      - run: zig build`

	return []byte(text)
}

//TODO: Zig_format to be implemented

// zigCmd represents the zig command
var zigCmd = &cobra.Command{
	Use:   "zig",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/workflows", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		fmt.Println("Creating cicd file for zig...")
		os.WriteFile(".github/workflows/zig.yml", zig(), 0644)
	},
}

func init() {
	rootCmd.AddCommand(zigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
