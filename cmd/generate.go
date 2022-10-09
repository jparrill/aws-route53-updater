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
		internal.Generator(ZoneID, Action, DNSRecordsFile, ChangeComment, OutputFormat, OutputPath, Filters...)

	},
}

func init() {
	rootCmd.AddCommand(generate)
	generate.PersistentFlags().StringVarP(&Action, "action", "a", "DELETE", "(UPPERCASE) Action to generate the BatchFile: (UPSERT|DELETE|CREATE)")
	generate.PersistentFlags().StringVarP(&DNSRecordsFile, "recordsfile", "r", "assets/samples/records.json", "AWS Route53 generated file from command 'aws route53 list-resource-record-sets ...'")
	generate.PersistentFlags().StringVarP(&ChangeComment, "comment", "c", "Change over Route53 platform in AWS", "Comment about the change to be submitted")
	generate.PersistentFlags().StringSliceVarP(&Filters, "recordTypes", "t", []string{}, "Filters or RecordTypes to just perform actions over them")
	generate.Flags().BoolP("info", "i", false, "Generates a batch file for AWS Route53 API with changes")
}
