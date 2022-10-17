package awsRoute53BG

import (
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func TestParseFileSuccess(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./importer_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./importer_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/../.."
	GenRightDNSRecordsFile := BasePath + "/" + dnsRecordsFile
	_, err = ParseFile(GenRightDNSRecordsFile)
	g.Expect(err).Should(BeNil())
}

func TestParseFileFailWrongPath(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./importer_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./importer_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/../.."
	GenRightDNSRecordsFile := BasePath + WrongPath
	_, err = ParseFile(GenRightDNSRecordsFile)
	g.Expect(err).Should(MatchError("Error loading RecordsFile, FileNotFound: \n - stat /Users/jparrill/RedHat/RedHat_Engineering/hypershift/repos/aws-route53-updater/pkg/awsRoute53BG/../../temepe/fail.json: no such file or directory"))
}

func TestParseFileFailWrongType(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./importer_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./importer_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/../.."
	GenRightDNSRecordsFile := BasePath + dnsRecordsYamlFile
	_, err = ParseFile(GenRightDNSRecordsFile)
	g.Expect(err).Should(MatchError("Filetype not implemented: \n - .txt"))
}
