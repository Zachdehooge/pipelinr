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

func golang() []byte {
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
          go build main.go

  gosec:
    name: Run Gosec Security Scan
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: "1.22"

      - name: Install Gosec
        run: go install github.com/securego/gosec/v2/cmd/gosec@latest

      - name: Run Gosec (fail only on HIGH severity)
        run: |
          gosec -fmt=json -severity=high ./... | tee gosec-report.json
          # Extract the number of HIGH severity issues
          HIGH_COUNT=$(jq '[.Issues[] | select(.severity == "HIGH")] | length' gosec-report.json)
          echo "Found $HIGH_COUNT high severity issues."
          if [ "$HIGH_COUNT" -gt 0 ]; then
            echo "High severity issues detected!"
            exit 1
          fi`

	return []byte(text)
}

func golang_format() []byte {
	var text = `name: Go Format

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
	Short: "starter pipeline for golang projects",
	Long:  `A starter pipeline with a formatter, builder, linter and security tests for your golang project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/workflows", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		color.Cyan("Creating cicd file for golang...")
		os.WriteFile(".github/workflows/golang.yml", golang(), 0644)

		color.Cyan("Creating formater file...")
		os.WriteFile(".github/workflows/golang_format.yml", golang_format(), 0644)

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
