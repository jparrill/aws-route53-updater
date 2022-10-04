package awsRoute53BG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type JSONFile struct {
	FilePath   string
	AWSRecords *AWSRecords
}

type YAMLFile struct {
	FilePath   string
	AWSRecords *AWSRecords
}

type TEXTFile struct {
	FilePath   string
	AWSRecords *AWSRecords
}

type Load interface {
	RecordsFile()
}

func (j *JSONFile) RecordsFile() {

	file, err := ioutil.ReadFile(j.FilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(file), &j.AWSRecords)
	if err != nil {
		panic(err)
	}
}

func (y *YAMLFile) RecordsFile() {}

func ParseFile(filePath string) AWSRecords {

	data := AWSRecords{}

	switch filepath.Ext(filePath) {
	case ".json":
		f := JSONFile{
			FilePath:   filePath,
			AWSRecords: &data,
		}

		f.RecordsFile()

	default:
		panic(fmt.Errorf("Filetype not implemented: %s\n", filepath.Ext(filePath)))
	}

	return data
}
