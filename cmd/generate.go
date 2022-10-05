package cmd

import (
	"github.com/jparrill/aws-route53-updater/internal"
	"github.com/spf13/cobra"
)

var generate = &cobra.Command{
	Use:   "generate",
	Short: "Generates a batch file for AWS Route53 API with changes",
	Long: `Generates a batch file for AWS Route53 API with changes. 
	
	Changes includes:
		Some Actions ("UPSERT"|"DELETE"|"CREATE")
		Over some resource Types: ("SOA"|"A"|"TXT"|"NS"|"CNAME"|"MX"|"NAPTR"|"PTR"|"SRV"|"SPF"|"AAAA"|"CAA"|"DS")`,

	Run: func(cmd *cobra.Command, args []string) {
		internal.Generator(ZoneID, Action, DNSRecordsFile, ChangeComment, OutputFormat, Filters...)

	},
}

func init() {
	rootCmd.AddCommand(generate)
	generate.Flags().BoolP("info", "i", false, "Generates a batch file for AWS Route53 API with changes")
}
