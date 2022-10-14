package internal

import (
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	// Wrong ones
	GenWrongZoneID            = "SICK"
	GenWrongAction            = "DILATE"
	GenWrongDNSRecordsFile    = "/temepe/sourceAWS.json"
	GenWrongOutputPath        = "/temepe/fail.json"
	GenEmptyRelDNSRecordsFile = "assets/samples/empty_records.json"
	GenWrongOutputFormat      = "jason"
	// Right ones
	GenRightZoneID            = "Z02718293M33QHDEQBROL"
	GenRightAction            = "DELETE"
	GenRightRelDNSRecordsFile = "assets/samples/records.json"
	GenRightOutputPath        = "/tmp/test.json"
	GenRightOutputFormat      = "json"
	GenSampleComment          = "This is a sample comment"
)

func TestGeneratorFailParsingRecFile(t *testing.T) {
	g := NewWithT(t)
	err := Generator(GenRightZoneID, GenRightAction, GenWrongDNSRecordsFile, GenSampleComment, GenRightOutputFormat, GenRightOutputPath)
	g.Expect(err).Should(MatchError("Error Parsing /temepe/sourceAWS.json RecordsFile: \n - Error loading RecordsFile, FileNotFound: \n - stat /temepe/sourceAWS.json: no such file or directory"))
}

func TestGeneratorFailWrongZone(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./generator_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./generator_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/.."
	GenRightDNSRecordsFile := BasePath + "/" + GenRightRelDNSRecordsFile
	err = Generator(GenWrongZoneID, GenRightAction, GenRightDNSRecordsFile, GenSampleComment, GenRightOutputFormat, GenRightOutputPath)
	g.Expect(err).ShouldNot(BeNil())
}

func TestGeneratorFailWrongOutFormat(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./generator_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./generator_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/.."
	GenRightDNSRecordsFile := BasePath + "/" + GenRightRelDNSRecordsFile
	err = Generator(GenRightZoneID, GenRightAction, GenRightDNSRecordsFile, GenSampleComment, GenWrongOutputFormat, GenRightOutputPath)
	g.Expect(err).Should(HaveOccurred(), "Output format not implemented")
}

func TestGeneratorFailWrongOutPath(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./generator_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./generator_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/.."
	GenRightDNSRecordsFile := BasePath + "/" + GenRightRelDNSRecordsFile
	err = Generator(GenRightZoneID, GenRightAction, GenRightDNSRecordsFile, GenSampleComment, GenWrongOutputFormat, GenRightOutputPath)
	g.Expect(err).Should(MatchError("Error classifying or exporting data received: \n - Error exporting data to struct: \n - Output format not implemented, (json|stdout) are the only methods available: \n - jason"))
}

func TestGeneratorFailWrongAction(t *testing.T) {
	g := NewWithT(t)
	// Recover sample filepath
	TestPath, err := filepath.Abs("./generator_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./generator_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/.."
	GenRightDNSRecordsFile := BasePath + "/" + GenRightRelDNSRecordsFile

	err = Generator(GenRightZoneID, GenWrongAction, GenRightDNSRecordsFile, GenSampleComment, GenRightOutputFormat, GenRightOutputPath)
	g.Expect(err).ShouldNot(BeNil())

}

func TestGeneratorFailEmptyRecFile(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./generator_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./generator_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/.."
	GenEmptyDNSRecordsFile := BasePath + "/" + GenEmptyRelDNSRecordsFile
	err = Generator(GenRightZoneID, GenRightAction, GenEmptyDNSRecordsFile, GenSampleComment, GenRightOutputFormat, GenRightOutputPath)
	g.Expect(err).Should(MatchError("There is no data to process, exiting...: \n - {DELETE []}"))
}
