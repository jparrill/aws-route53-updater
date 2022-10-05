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
The base documentation is located here https://docs.aws.amazon.com/cli/latest/reference/route53/change-resource-record-sets.html.

The source PKG is located at github.com/jparrill/aws-route53-updater/pkg/awsRoute53BG which contains all the functions and documentation
for their consumption and use.`,
}

var (
	Action         string
	ZoneID         string
	Filters        []string
	DNSRecordsFile string
	ChangeComment  string
	OutputFormat   string
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("info", "i", false, "Basic info message")
	rootCmd.PersistentFlags().StringVarP(&Action, "action", "a", "DELETE", "(UPPERCASE) Action to generate the BatchFile: (UPSERT|DELETE|CREATE)")
	rootCmd.PersistentFlags().StringVarP(&ZoneID, "zoneid", "z", "Z02718293M33QHDEQBROL", "AWS Route53 ZoneID to be modified")
	rootCmd.PersistentFlags().StringVarP(&DNSRecordsFile, "recordsfile", "r", "assets/samples/records.json", "AWS Route53 generated file from command 'aws route53 list-resource-record-sets ...'")
	rootCmd.PersistentFlags().StringVarP(&ChangeComment, "comment", "c", "Change over Route53 platform in AWS", "Comment about the change to be submitted")
	rootCmd.PersistentFlags().StringVarP(&OutputFormat, "output", "o", "json", "Output Format for the file to be submitted to Route53 API: (json|yaml)")
	rootCmd.PersistentFlags().StringSliceVarP(&Filters, "filters", "f", []string{}, "Filters to just perform actions over them")

}
