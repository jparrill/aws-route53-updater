package internal

import (
	"fmt"

	"github.com/jparrill/aws-route53-updater/pkg/awsRoute53BG"
)

func Generator(zoneID, action, dnsRecordsFile, changeComment, outputFormat, outputPath string, filters ...string) {

	// Initialize the basic variables
	xChanges := make([]awsRoute53BG.Changes, 0, 0)

	// Recover and parse AWS Records file for Zone X
	DRecords := awsRoute53BG.ParseFile(dnsRecordsFile)
	DChanges := awsRoute53BG.ProcessData(action, DRecords.ResourceRecordSets, filters...)
	xChanges = append(xChanges, DChanges)

	AWSCFile := awsRoute53BG.ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", changeComment, zoneID),
		Changes: xChanges,
	}

	// Output
	switch outputFormat {
	case "stdout":
		awsRoute53BG.Exporter(outputFormat, AWSCFile, "")
	case "json":
		awsRoute53BG.Exporter(outputFormat, AWSCFile, outputPath)
	default:
		panic(fmt.Errorf("Exporter method not implemented: %s\n", outputFormat))
	}
}
