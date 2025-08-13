/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func copyYAMLPython(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "starter pipeline for python",
	Long:  `A starter pipeline with a formatter, builder, linter and security tests for your python project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/templates", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		color.Blue("Copying python.yml workflow...")
		if err := copyYAMLPython("./cmd/templates/python.yml", ".github/templates/python.yml"); err != nil {
			fmt.Printf("Error copying python.yml: %v\n", err)
		}

		color.Blue("Copying python_format.yml workflow...")
		if err := copyYAMLPython("./cmd/templates/python_format.yml", ".github/templates/python_format.yml"); err != nil {
			fmt.Printf("Error copying python_format.yml: %v\n", err)
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
