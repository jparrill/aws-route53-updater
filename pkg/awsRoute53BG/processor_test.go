package awsRoute53BG

import (
	"fmt"
	"path/filepath"
	"testing"

	. "github.com/onsi/gomega"
)

func TestProcessSuccessActionWithFilters(t *testing.T) {
	g := NewWithT(t)
	filters := make([]string, 0, 0)
	filters = append(filters, filterTXT)
	TestPath, err := filepath.Abs("./importer_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./importer_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/../.."
	GenRightDNSRecordsFile := BasePath + "/" + dnsRecordsFile
	DRecords, err := ParseFile(GenRightDNSRecordsFile)
	g.Expect(err).Should(BeNil())
	DChanges := ProcessData("DELETE", DRecords.ResourceRecordSets, filters...)
	g.Expect(DChanges.Action).Should(Equal("DELETE"))
}

func TestProcessSuccessDataWithFilters(t *testing.T) {
	g := NewWithT(t)
	filters := make([]string, 0, 0)
	filters = append(filters, filterTXT)
	TestPath, err := filepath.Abs("./importer_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./importer_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/../.."
	GenRightDNSRecordsFile := BasePath + "/" + dnsRecordsFile
	DRecords, err := ParseFile(GenRightDNSRecordsFile)
	g.Expect(err).Should(BeNil())
	DChanges := ProcessData("DELETE", DRecords.ResourceRecordSets, filters...)
	g.Expect(DChanges.Action).NotTo(ContainSubstring("TXT"))
}

func TestProcessSuccessDataNoFilters(t *testing.T) {
	g := NewWithT(t)
	TestPath, err := filepath.Abs("./importer_test.go")
	if err != nil {
		panic(fmt.Errorf("File ./importer_test.go Not found in local folder"))
	}
	BasePath := filepath.Dir(TestPath) + "/../.."
	GenRightDNSRecordsFile := BasePath + "/" + dnsRecordsFile
	DRecords, err := ParseFile(GenRightDNSRecordsFile)
	g.Expect(err).Should(BeNil())
	DChanges := ProcessData("DELETE", DRecords.ResourceRecordSets)
	g.Expect(DChanges.ResourceRecordSets).ShouldNot(BeEmpty())
}
