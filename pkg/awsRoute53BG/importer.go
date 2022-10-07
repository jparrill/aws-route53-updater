package awsRoute53BG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Load interface {
	RecordsFile() AWSRecords
}

func (j *JSONFile) RecordsFile() AWSRecords {

	var err error
	awsRec := AWSRecords{}
	RRS := make([]ResourceRecordSets, 0, 0)
	awsRec.ResourceRecordSets = RRS

	j.Data, err = ioutil.ReadFile(j.FilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(j.Data, &awsRec)
	if err != nil {
		panic(err)
	}

	return awsRec
}

func (y *YAMLFile) RecordsFile() {}

func ParseFile(filePath string) AWSRecords {

	var data []byte

	switch filepath.Ext(filePath) {
	case ".json":
		f := JSONFile{
			FilePath: filePath,
			Data:     data,
		}

		return f.RecordsFile()

	default:
		panic(fmt.Errorf("Filetype not implemented: %s\n", filepath.Ext(filePath)))
	}
}
