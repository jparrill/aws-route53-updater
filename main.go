package main

import (
	"html/template"
	"os"
)

type Record struct {
	Action    string
	DNSRecord string
	Type      string
}

type JSONFile string
type YAMLFile string
type CSVFile string

type Load interface {
	RecordsFile()
}

func (j JSONFile) RecordsFile() {}
func (y YAMLFile) RecordsFile() {}
func (c CSVFile) RecordsFile()  {}

func parseFile(path string) ([]Record, error) {}

func main() {
	records := make([]Record, 0, 0)

	drecord := Record{
		Action:    "DELETE",
		DNSRecord: "ovn-sbdb-jparrill-hosted-external-dns.jparrill-hosted.aws.kerbeross.com.",
		Type:      "TXT",
	}

	t, err := template.New("r53-records").ParseFiles("assets/templates/batch.tmpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, drecord)
	if err != nil {
		panic(err)
	}

}

func storeEntry(entry Record, xRecord *[]Record) error {}
