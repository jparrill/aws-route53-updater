package awsRoute53BG

import (
	"bytes"
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

func Exporter(format string, data ChangeJson, path string) {

	switch format {
	case "stdout":
		b, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		j := JSONFile{
			Format:   format,
			FilePath: path,
			Data:     b,
		}

		fmt.Println(string(j.Data))

	case "json":
		b, err := json.Marshal(data)
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
