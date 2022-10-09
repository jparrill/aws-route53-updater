package cmd

import (
	"github.com/jparrill/aws-route53-updater/internal"
	"github.com/spf13/cobra"
)

var recovery = &cobra.Command{
	Use:   "recovery",
	Short: "Generates file from AWS Route53 API with the DNS records for a concrete Zone ID",
	Long: `Generates file from AWS Route53 API with the DNS records for a concrete Zone ID.

	This is the file you will need to use in Generate subcommand to filter the entries and execute actions.
	`,

	Run: func(cmd *cobra.Command, args []string) {
		internal.Recover(ZoneID, OutputPath, OutputFormat, Filters...)
	},
}

func init() {
	rootCmd.AddCommand(recovery)
	recovery.Flags().BoolP("info", "i", false, "Generates file from AWS Route53 API with the DNS records for a concrete Zone ID")
}
