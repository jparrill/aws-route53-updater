package main

import (
	"fmt"

	"github.com/jparrill/pkg/awsRoute53BG"
)

func main() {
	// Initialize the basic variables
	ZoneID := "Z02718293M33QHDEQBROL"
	xChanges := make([]awsRoute53BG.Changes, 0, 0)
	filters := make([]string, 0, 0)
	filters = append(filters, "TXT")

	// Recover and parse AWS Records file for Zone X
	DRecords := awsRoute53BG.ParseFile("assets/samples/records.json")
	DChanges := awsRoute53BG.ProcessData("DELETE", DRecords.ResourceRecordSets, filters...)
	xChanges = append(xChanges, DChanges)

	AWSCFile := awsRoute53BG.ChangeJson{
		Comment: fmt.Sprintf("Changes over Route53 AWS in Zone: %v", ZoneID),
		Changes: xChanges,
	}

	// Output
	awsRoute53BG.Exporter("stdout", AWSCFile)
}
