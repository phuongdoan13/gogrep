/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "Search for matching pattern in a file",
	Long: `
		Grep is a useful command to search for a pattern in a file, which is short for 'global regular expression print'. 
		It searches for the pattern in the file and prints the line that contains the pattern.

					grep [options] pattern [file...]
		
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grep called")
	},
}

func init() {
	rootCmd.AddCommand(grepCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grepCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grepCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
