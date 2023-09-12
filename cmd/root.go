/*
Copyright © 2023 deepakchethan <deepakchethan@outlook.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version bool
var quiet bool
var verbose bool
var help bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lnkdocs",
	Short: "Lnkdocs like mkDocs for bookmarks based page",
	Long:  `Lnkdocs like mkDocs for bookmarks based page`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "Show the version and exit")
	rootCmd.Flags().BoolVarP(&quiet, "quiet", "q", false, "Silence warnings")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().BoolVarP(&help, "help", "h", false, "Show this message and exit")
}
