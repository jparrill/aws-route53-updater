package internal

import (
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	// Wrong ones
	GenWrongZoneID         = "SICK"
	GenWrongAction         = "DILATE"
	GenWrongDNSRecordsFile = "/temepe/sourceAWS.json"
	GenWrongOutputPath     = "/temepe/fail.json"
	GenWrongOutputFormat   = "jason"
	// Right ones
	GenRightZoneID            = "Z02718293M33QHDEQBROL"
	GenRightAction            = "DELETE"
	GenRightRelDNSRecordsFile = "assets/samples/records.json"
	GenRightOutputPath        = "/tmp/test.json"
	GenRightOutputFormat      = "json"
	GenSampleComment          = "This is a sample comment"
)

func init() {
}

func TestGeneratorFailParsingRecFile(t *testing.T) {
	g := NewWithT(t)
	err := Generator(GenWrongZoneID, GenRightAction, GenWrongDNSRecordsFile, GenSampleComment, GenRightOutputFormat, GenRightOutputPath)
	g.Expect(err).Should(MatchError("Error Parsing /temepe/sourceAWS.json RecordsFile: \n - Error loading RecordsFile, FileNotFound: \n - stat /temepe/sourceAWS.json: no such file or directory"))
}

func TestGeneratorFailWrongZone(t *testing.T) {
	g := NewWithT(t)
	err := Generator(GenWrongZoneID, GenRightAction, GenRightRelDNSRecordsFile, GenSampleComment, GenRightOutputFormat, GenRightOutputPath)
	g.Expect(err).ShouldNot(BeNil())
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
