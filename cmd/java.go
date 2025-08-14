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

func copyYAMLJava(src, dst string) error {
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

// javaCmd represents the java command
var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "starter pipeline for java",
	Long:  `A starter pipeline with a formatter, builder, linter and security tests for your java project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/templates", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		color.Red("Copying java.yml workflow...")
		if err := copyYAMLJava("./cmd/templates/java.yml", ".github/workflows/java.yml"); err != nil {
			fmt.Printf("Error copying java.yml: %v\n", err)
		}

		color.Red("Copying java_format.yml workflow...")
		if err := copyYAMLJava("./cmd/templates/java_format.yml", ".github/workflows/java_format.yml"); err != nil {
			fmt.Printf("Error copying java_format.yml: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(javaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// javaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// javaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
