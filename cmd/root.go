/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"strings"
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
		fmt.Println("grep called")
		pattern := args[0]
		fileName := args[1]
		
		if config.IsIgnoreCase {
			pattern = strings.ToLower(pattern)
		}
		
		pkg.Grep(pattern, fileName)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gogrep.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&config.IsIgnoreCase, "ignore-case", "i", false, "Ignore case distinctions in both the PATTERN and the input files.")
}


