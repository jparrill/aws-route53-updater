package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type ChangeJson struct {
	Comment string    `json:"Comment"`
	Changes []Changes `json:"Changes"`
}

type Changes struct {
	Action             string               `json:"Action"`
	ResourceRecordSets []ResourceRecordSets `json:"ResourceRecordSets"`
}
type AWSRecords struct {
	ResourceRecordSets []ResourceRecordSets `json:"ResourceRecordSets"`
}
type ResourceRecords struct {
	Value string `json:"Value"`
}
type ResourceRecordSets struct {
	Name            string            `json:"Name"`
	Type            string            `json:"Type"`
	TTL             int               `json:"TTL"`
	ResourceRecords []ResourceRecords `json:"ResourceRecords"`
}

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

func parseFile(filePath string) AWSRecords {

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

func processData(data []ResourceRecordSets) ChangeJson {
	changedData := ChangeJson{
		Comment: "DNS Update from AWS Changer",
	}

	fmt.Println(data)

	return changedData
}

//func generateChangeFile(outPath string, data AWSRecords) error {}

func main() {

	DRecords := parseFile("assets/samples/records.json")
	CDRecords := processData(DRecords.ResourceRecordSets)
	fmt.Println(CDRecords)
	//generateChangeFile("out/AWSRecords.json")

}

//func storeEntry(entry Record, xRecord *[]Record) error {}
