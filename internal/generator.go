package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/jparrill/aws-route53-updater/pkg/awsRoute53BG"
)

func Generator(zoneID, action, dnsRecordsFile, changeComment, outputFormat, outputPath string, filters ...string) error {

	// Initialize the basic variables
	xChanges := make([]awsRoute53BG.Changes, 0, 0)

	// Recover and parse AWS Records file for Zone X
	DRecords, err := awsRoute53BG.ParseFile(dnsRecordsFile)
	if err != nil {
		return fmt.Errorf("Error Parsing %s RecordsFile: \n - %v", dnsRecordsFile, err)
	}
	DChanges := awsRoute53BG.ProcessData(action, DRecords.ResourceRecordSets, filters...)
	if len(DChanges.ResourceRecordSets) <= 0 {
		return fmt.Errorf("There is no data to process, exiting...: \n - %v", DChanges)
	}
	xChanges = append(xChanges, DChanges)

	var buff bytes.Buffer
	kind := "gen"
	enc := gob.NewEncoder(&buff)
	err = enc.Encode(awsRoute53BG.ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", changeComment, zoneID),
		Changes: xChanges,
	})
	if err != nil {
		return fmt.Errorf("Error storing data for processing into buffer: \n - %v", err)
	}

	// Output
	err = awsRoute53BG.Classifier(outputFormat, &buff, outputPath, kind)
	if err != nil {
		return fmt.Errorf("Error classifying or exporting data received: \n - %v", err)
	}

	return nil
}
