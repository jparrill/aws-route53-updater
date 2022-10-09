package awsRoute53BG

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Format interface {
	Export()
}

func (j JSONFile) Export() {
	fmt.Println("Exporting data to JSONFile: ", j.FilePath)
	var indentedBytes bytes.Buffer

	err := json.Indent(&indentedBytes, j.Data, "", "  ")
	if err != nil {
		panic(err)
	}

	_ = ioutil.WriteFile(j.FilePath, indentedBytes.Bytes(), 0644)

}

func (y YAMLFile) Export() {}

func Classifier(format string, rawData *bytes.Buffer, path string, kind string) {

	dec := gob.NewDecoder(rawData)

	switch kind {
	case "gen":
		var AWSChange ChangeJson
		dec.Decode(&AWSChange)
		AWSChange.Exporter(format, path, kind)

	case "rec":
		var RRS localroute53RRS
		dec.Decode(&RRS)
		RRS.Exporter(format, path, kind)

	default:
		panic(fmt.Errorf("Kind type not implemented: %s\n", kind))
	}

}

func (rrs localroute53RRS) Exporter(format string, path string, kind string) {

	switch format {
	case "stdout":
		b, err := json.Marshal(rrs)
		if err != nil {
			panic(err)
		}
		j := JSONFile{
			Format:   format,
			FilePath: "",
			Data:     b,
		}

		fmt.Println(string(j.Data))

	case "json":
		b, err := json.Marshal(rrs)
		if err != nil {
			panic(err)
		}
		j := JSONFile{
			Format:   format,
			FilePath: path,
			Data:     b,
		}

		j.Export()

	default:
		panic(fmt.Errorf("Output method not implemented: %s\n", format))
	}
}

func (c *ChangeJson) Exporter(format string, path string, kind string) {

	switch format {
	case "stdout":
		b, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}
		j := JSONFile{
			Format:   format,
			FilePath: "",
			Data:     b,
		}

		fmt.Println(string(j.Data))

	case "json":
		b, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}
		j := JSONFile{
			Format:   format,
			FilePath: path,
			Data:     b,
		}

		j.Export()

	default:
		panic(fmt.Errorf("Output method not implemented: %s\n", format))
	}
}
