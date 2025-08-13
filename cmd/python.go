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

func python() []byte {
	var text = `name: Python CI/CD

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
  run:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: '3.13'
      - run: python main.py


  bandit:
    name: Bandit Security Scan
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Python
        uses: actions/setup-python@v5
        with:
          python-version: "3.11"

      - name: Install Bandit and jq
        run: |
          pip install bandit
          sudo apt-get update && sudo apt-get install -y jq

      - name: Run Bandit (fail only on HIGH severity)
        run: |
          bandit -r . -f json -o bandit-report.json || true
          HIGH_COUNT=$(jq '[.results[] | select(.issue_severity == "HIGH")] | length' bandit-report.json)
          echo "Found $HIGH_COUNT high severity issues."
          if [ "$HIGH_COUNT" -gt 0 ]; then
            echo "High severity issues detected! Failing the build."
            exit 1
          fi`

	return []byte(text)
}

func python_format() []byte {
	var text = `

name: Python Format

on:
  pull_request:
    branches: [ main ]
  push:
    branches:
      - main
      - temp_pipe
  workflow_dispatch:

jobs:
  format:
    runs-on: ubuntu-latest
    permissions:
      contents: write
    
    steps:
    - uses: actions/checkout@v3
      with:
        ref: ${{ github.head_ref }}

    - name: Set up Python
      uses: actions/setup-python@v4
      with:
        python-version: '3.12'

    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        pip install black

    - name: Format with Black
      run: |
        black main.py
      

    - name: Commit changes
      uses: stefanzweifel/git-auto-commit-action@v4
      with:
        commit_message: "Apply Black formatting"
        file_pattern: |
          ./main.py`

	return []byte(text)
}

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
		color.Blue("Creating cicd file for python...")
		os.WriteFile(".github/workflows/python.yml", python(), 0644)

		color.Blue("Creating formater file for python...")
		os.WriteFile(".github/workflows/python_format.yml", python_format(), 0644)
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
