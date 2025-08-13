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

func zig() []byte {
	var text = `name: CI/CD

on:
  push:
    branches:
      - main
      - temp_pipe
  pull_request:
    branches:
      - main
  workflow_dispatch:
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

func zig_format() []byte {
	var text = `name: Format

on:
  push:
    branches:
      - main
      - temp_pipe
  pull_request:
    branches:
      - main
  workflow_dispatch:

jobs:
  format:
    runs-on: ubuntu-latest
    name: Format
    steps:
      - uses: actions/checkout@v3
      - uses: mlugg/setup-zig@v2
      - run: zig fmt src/*.zig


# Commit changes if any
      - name: Check for changes
        id: git-check
        run: |
          git status --porcelain
          echo "changes=$(git status --porcelain | wc -l)" >> $GITHUB_OUTPUT

      - name: Commit changes
        if: steps.git-check.outputs.changes > 0
        run: |
          git config --local user.email "github-actions[bot]@users.noreply.github.com"
          git config --local user.name "github-actions[bot]"
          git add -A
          git commit -m "Auto-format Go code"
          git push`

	return []byte(text)
}

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
		color.Yellow("Creating cicd file for zig...")
		os.WriteFile(".github/workflows/zig.yml", zig(), 0644)

		color.Yellow("Creating format file for zig...")
		os.WriteFile(".github/workflows/zig_format.yml", zig_format(), 0644)
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
