package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aru",
	Short: "AWS Route53 Changes Generator",
	Long: `AWS Route53 Changes batch file Generator.

	This program generates the needed file to update the Route53 DNS entries focused in a concrete Hosted Zone.
	The base documentation is located here https://docs.aws.amazon.com/cli/latest/reference/route53/change-resource-record-sets.html`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Cobra Flags
	// - Action
	// - ZoneID
	// - Filters
	// - DNS Records FilePath
	// - Comment
	// - OutFormat

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
