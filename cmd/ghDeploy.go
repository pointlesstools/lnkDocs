/*
Copyright © 2023 deepakchethan <deepak.chethan.s@sap.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ghDeployCmd represents the ghDeploy command
var ghDeployCmd = &cobra.Command{
	Use:   "gh-deploy",
	Short: "Deploy your documentation to GitHub Pages",
	Long:  `Deploy your documentation to GitHub Pages`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ghDeploy called")
	},
}

func init() {
	rootCmd.AddCommand(ghDeployCmd)
}
