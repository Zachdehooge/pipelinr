/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func mainFile() []byte {
	var text = `
name: CI/CD

on:
  push:
    branches:
      - main
  pull_request:
    branches:
      - main
jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '>=1.23.0'
      - name: Build
        run: |
          echo "Running build..."
  
          go mod tidy
          go build main.go`

	return []byte(text)
}

func formatFile() []byte {
	var text = `
name: Go Format

on:
  pull_request:
    branches: [ main ]
  push:
    branches: [ main ]

jobs:
  format:
    name: Format Go Code
    runs-on: ubuntu-latest
    permissions:
      contents: write
      pull-requests: write

    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          ref: ${{ github.head_ref }}

      # Setup Go
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '>=1.23'

      # Run Go formatters
      - name: Format Go code
        run: |
          go install golang.org/x/tools/cmd/goimports@latest
          go fmt ./...
          goimports -w .

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

var golangCmd = &cobra.Command{
	Use:   "golang",
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
		fmt.Println("Creating GitHub Actions file for Golang...")
		os.WriteFile(".github/workflows/golang.yml", mainFile(), 0644)

		fmt.Println("Creating format on pull request...")
		os.WriteFile(".github/workflows/format.yml", formatFile(), 0644)

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
