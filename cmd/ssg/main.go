package main

import (
	"fmt"
	"os"

	"github.com/deelawn/ssg/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ssg",
	Short: "A simple static site generator",
	Long:  `SSG is a CLI tool for generating static websites from markdown files with theme support.`,
}

func init() {
	rootCmd.AddCommand(commands.InitCmd)
	rootCmd.AddCommand(commands.BuildCmd)
	rootCmd.AddCommand(commands.AddCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
