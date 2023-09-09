/*
Copyright © 2023 Harry Doan

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/phuongdoan13/gogrep/pkg"
	"github.com/phuongdoan13/gogrep/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	RunE: func(cmd *cobra.Command, args []string) (error) {
		pattern := args[0]
		fileName := args[1]
		
		result := pkg.Grep(pattern, fileName)
		ans := formatOutput(result)

		fmt.Fprintf(cmd.OutOrStdout(), ans)
		return nil
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
	rootCmd.Flags().BoolP(config.IgnoreCaseFlag, "i", false, "Ignore case distinctions in both the PATTERN and the input files.")
	rootCmd.Flags().BoolP(config.LineNumberFlag, "n", false, "Prefix each line of the matching output with the line number in the input file.")
	
	viper.BindPFlag(config.IgnoreCaseFlag, rootCmd.Flags().Lookup(config.IgnoreCaseFlag))
	viper.BindPFlag(config.LineNumberFlag, rootCmd.Flags().Lookup(config.LineNumberFlag))
}

func formatOutput(result []pkg.PairLineNumberAndLine) string {
	var ans string

	for _, pair := range result {
		if viper.GetBool(config.LineNumberFlag) {
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