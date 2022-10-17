package awsRoute53BG

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Load interface {
	RecordsFile() AWSRecords
}

func (j *JSONFile) RecordsFile(awsRec AWSRecords) (AWSRecords, error) {

	var err error
	RRS := make([]ResourceRecordSets, 0, 0)
	awsRec.ResourceRecordSets = RRS

	if _, err := os.Stat(j.FilePath); errors.Is(err, os.ErrNotExist) {
		return awsRec, fmt.Errorf("Error loading RecordsFile, FileNotFound: \n - %v", err)
	}
	j.Data, err = ioutil.ReadFile(j.FilePath)

	err = json.Unmarshal(j.Data, &awsRec)
	if err != nil {
		return awsRec, fmt.Errorf("Error unmarshaling data from RecordsFile: \n - %v", err)
	}

	return awsRec, nil
}

func ParseFile(filePath string) (AWSRecords, error) {

	var data []byte
	awsRec := AWSRecords{}

	switch filepath.Ext(filePath) {
	case ".json":
		f := JSONFile{
			FilePath: filePath,
			Data:     data,
		}

		return f.RecordsFile(awsRec)

	default:
		return awsRec, fmt.Errorf("Filetype not implemented: \n - %s", filepath.Ext(filePath))
	}
}
