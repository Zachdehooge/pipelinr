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

func copyYAMLZig(src, dst string) error {
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

// zigCmd represents the zig command
var zigCmd = &cobra.Command{
	Use:   "zig",
	Short: "starter pipeline for zig",
	Long:  `A starter pipeline with a formatter, builder, linter and security tests for your zig project`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll(".github/templates", 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}
		color.Yellow("Copying zig.yml workflow...")
		if err := copyYAMLZig("./cmd/templates/zig.yml", ".github/workflows/zig.yml"); err != nil {
			fmt.Printf("Error copying zig.yml: %v\n", err)
		}

		color.Yellow("Copying zig_format.yml workflow...")
		if err := copyYAMLZig("./cmd/templates/zig_format.yml", ".github/workflows/zig_format.yml"); err != nil {
			fmt.Printf("Error copying zig_format.yml: %v\n", err)
		}
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
