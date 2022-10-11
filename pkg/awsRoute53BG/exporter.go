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

func (j JSONFile) Export() error {
	fmt.Println("Exporting data to JSONFile: ", j.FilePath)
	var indentedBytes bytes.Buffer

	err := json.Indent(&indentedBytes, j.Data, "", "  ")
	if err != nil {
		return fmt.Errorf("Error indenting data into struct: %v\n", err)
	}

	err = ioutil.WriteFile(j.FilePath, indentedBytes.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("Error writing data into file %s: %v\n", j.FilePath, err)
	}

	return nil

}

func (y YAMLFile) Export() {}

func Classifier(format string, rawData *bytes.Buffer, path string, kind string) error {

	dec := gob.NewDecoder(rawData)

	switch kind {
	case "gen":
		var AWSChange ChangeJson

		err := dec.Decode(&AWSChange)
		if err != nil {
			return fmt.Errorf("Error decoding data from buffer: %v\n", err)
		}

		err = Exporter(format, path, kind, AWSChange)
		if err != nil {
			return fmt.Errorf("Error exporting data to struct: %v\n", err)
		}

	case "rec":
		var RRS localroute53RRS

		err := dec.Decode(&RRS)
		if err != nil {
			return fmt.Errorf("Error decoding data from buffer: %v\n", err)
		}

		Exporter(format, path, kind, RRS)

	default:
		return fmt.Errorf("Kind type not implemented: %s\n", kind)
	}

	return nil
}

func Exporter(format string, path string, kind string, component interface{}) error {

	switch component.(type) {
	case localroute53RRS, ChangeJson:
		b, err := json.Marshal(component)
		if err != nil {
			return fmt.Errorf("Error marshaling output JSON: %v\n", err)
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

			err := j.Export()
			if err != nil {
				return fmt.Errorf("Error exporting to JSON: %v\n", err)
			}

		} else {
			return fmt.Errorf("Output method not implemented, (json|stdout) methods implemented): %s\n", format)
		}

	default:
		return fmt.Errorf("Component not implemented: %s\n", format)
	}

	return nil
}
