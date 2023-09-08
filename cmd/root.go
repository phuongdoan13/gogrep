/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/phuongdoan13/gogrep/pkg"
	"github.com/phuongdoan13/gogrep/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "Search for matching pattern in a file",
	Args:  cobra.ExactArgs(2),
	Long: `
		Grep is a useful command to search for a pattern in a file, which is short for 'global regular expression print'. 
		It searches for the pattern in the file and prints the line that contains the pattern.

					grep [options] pattern [file...]
		
	`,
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]
		fileName := args[1]
		
		result := pkg.Grep(pattern, fileName)
		ans := formatOutput(result)

		fmt.Fprintf(cmd.OutOrStdout(), ans)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&config.IsIgnoreCase, "ignore-case", "i", false, "Ignore case distinctions in both the PATTERN and the input files.")
	rootCmd.Flags().BoolVarP(&config.IsPrintLnWithNumLine, "line-number", "n", false, "Prefix each line of the matching output with the line number in the input file.")
}

func formatOutput(result []pkg.PairLineNumberAndLine) string {
	var ans string

	for _, pair := range result {
		if config.IsPrintLnWithNumLine {
			ans += fmt.Sprintf("%d %s\n", pair.LineNumber, pair.Line)
		} else {
			ans += fmt.Sprintf("%s\n", pair.Line)
		}
	}

	return ans
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}