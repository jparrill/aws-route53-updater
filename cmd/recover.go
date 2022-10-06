package cmd

import (
	"encoding/json"
	"fmt"

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
		rrs := internal.RecoverRecordSet(ZoneID, DNSRecordsFile, OutputFormat, Filters...)
		b, _ := json.Marshal(rrs)
		fmt.Println(string(b))

		//awsRoute53BG.Exporter("stdout", rrs)
	},
}

func init() {
	rootCmd.AddCommand(recovery)
	recovery.PersistentFlags().StringVarP(&ZoneID, "zoneid", "z", "Z02718293M33QHDEQBROL", "AWS Route53 ZoneID to be modified")
	recovery.PersistentFlags().StringVarP(&DNSRecordsFile, "recordsfile", "r", "assets/samples/records.json", "AWS Route53 generated file from command 'aws route53 list-resource-record-sets ...'")
	recovery.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "json", "Output Format for the file to be submitted to Route53 API: (json|yaml)")
	recovery.PersistentFlags().StringSliceVarP(&Filters, "filters", "f", []string{}, "Filters to just perform actions over them")

	recovery.Flags().BoolP("info", "i", false, "Generates file from AWS Route53 API with the DNS records for a concrete Zone ID")
}
