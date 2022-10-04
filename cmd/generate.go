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

var (
	Action         string
	ZoneID         string
	Filters        []string
	DNSRecordsFile string
	ChangeComment  string
	OutputFormat   string
)

func init() {
	rootCmd.AddCommand(generate)
	rootCmd.PersistentFlags().StringVarP(&Action, "action", "a", "DELETE", "(UPPERCASE) Action to generate the BatchFile: (UPSERT|DELETE|CREATE)")
	rootCmd.PersistentFlags().StringVarP(&ZoneID, "zoneid", "z", "Z02718293M33QHDEQBROL", "AWS Route53 ZoneID to be modified")
	rootCmd.PersistentFlags().StringVarP(&DNSRecordsFile, "recordsfile", "r", "assets/samples/records.json", "AWS Route53 generated file from command 'aws route53 list-resource-record-sets ...'")
	rootCmd.PersistentFlags().StringVarP(&ChangeComment, "comment", "c", "Change over Route53 platform in AWS", "Comment about the change to be submitted")
	rootCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "json", "Output Format for the file to be submitted to Route53 API: (json|yaml)")
	rootCmd.PersistentFlags().StringSliceVarP(&Filters, "filters", "f", []string{}, "Filters to just perform actions over them")
	generate.Flags().BoolP("info", "i", false, "Basic info message")
}