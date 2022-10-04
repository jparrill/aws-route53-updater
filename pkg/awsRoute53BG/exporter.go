package awsRoute53BG

import (
	"encoding/json"
	"fmt"
)

func Exporter(output string, data ChangeJson) {

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
