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
		Exporter(format, path, kind, AWSChange)

	case "rec":
		var RRS localroute53RRS
		dec.Decode(&RRS)
		Exporter(format, path, kind, RRS)

	default:
		panic(fmt.Errorf("Kind type not implemented: %s\n", kind))
	}

}

func Exporter(format string, path string, kind string, component interface{}) {

	switch component.(type) {
	case localroute53RRS, ChangeJson:
		b, err := json.Marshal(component)
		if err != nil {
			panic(err)
		}
		j := JSONFile{
			Format:   format,
			FilePath: path,
			Data:     b,
		}

		if format == "stdout" {

			j.FilePath = ""
			fmt.Println(string(j.Data))

		} else if format == "json" {

			j.Export()

		} else {
			panic(fmt.Errorf("Output method not implemented, (json|stdout) methods implemented): %s\n", format))
		}

	default:
		panic(fmt.Errorf("Component not implemented: %s\n", format))
	}
}
