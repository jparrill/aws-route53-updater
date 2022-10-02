package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/exp/slices"
)

type ChangeJson struct {
	Comment string    `json:"Comment"`
	Changes []Changes `json:"Changes"`
}

type Changes struct {
	Action             string                `json:"Action"`
	ResourceRecordSets []CResourceRecordSets `json:"ResourceRecordSets"`
}
type AWSRecords struct {
	ResourceRecordSets []ResourceRecordSets `json:"ResourceRecordSets"`
}
type ResourceRecords struct {
	Value string `json:"Value"`
}
type CResourceRecordSets struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
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

func processData(action string, data []ResourceRecordSets, filters ...string) Changes {
	newData := make([]CResourceRecordSets, 0, 0)

	if len(filters) == 0 {
		for _, v := range data {
			if v.Type != "SOA" && v.Type != "NS" {
				newEntry := CResourceRecordSets{
					Name: v.Name,
					Type: v.Type,
				}
				newData = append(newData, newEntry)

			}
		}
	} else {
		for _, v := range data {
			if v.Type != "SOA" && v.Type != "NS" {
				if slices.Contains(filters, v.Type) {
					newEntry := CResourceRecordSets{
						Name: v.Name,
						Type: v.Type,
					}
					newData = append(newData, newEntry)
				}
			}
		}
	}

	c := Changes{
		Action:             action,
		ResourceRecordSets: newData,
	}

	return c
}

func renderJson(output string, data ChangeJson) {

	switch output {
	case "stdout":
		b, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))

	default:
		panic(fmt.Errorf("Output method not implemented: %s\n", output))
	}
}

func main() {

	// Initialize the basic variables
	ZoneID := "Z02718293M33QHDEQBROL"
	xChanges := make([]Changes, 0, 0)
	filters := make([]string, 0, 0)
	filters = append(filters, "TXT")

	// Recover and parse AWS Records file for Zone X
	DRecords := parseFile("assets/samples/records.json")
	DChanges := processData("DELETE", DRecords.ResourceRecordSets, filters...)
	xChanges = append(xChanges, DChanges)

	AWSCFile := ChangeJson{
		Comment: fmt.Sprintf("Changes over Route53 AWS in Zone: %v", ZoneID),
		Changes: xChanges,
	}

	// Output
	renderJson("stdout", AWSCFile)
}
