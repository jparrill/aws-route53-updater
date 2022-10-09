package internal

import (
	"bytes"
	"encoding/gob"
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

	var buff bytes.Buffer
	kind := "gen"
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(awsRoute53BG.ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", changeComment, zoneID),
		Changes: xChanges,
	})
	if err != nil {
		panic(err)
	}

	// Output
	awsRoute53BG.Classifier(outputFormat, &buff, outputPath, kind)
}
