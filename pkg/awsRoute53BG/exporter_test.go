package awsRoute53BG

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

const (
	WrongFormat        = "jason"
	WrongRawData       = "{[{}]}"
	WrongPath          = "/temepe/fail.json"
	WrongKind          = "rev"
	WrongAction        = "DILATE"
	WrongZoneID        = "Z02718293M33EQBROCOLY"
	RightFormat        = "json"
	RightAction        = "DELETE"
	RightPath          = "/tmp/test.json"
	RightRawData       = `{"ResourceRecordSets":[]}`
	RecRightKind       = "rec"
	GenRightKind       = "gen"
	RightZoneID        = "Z02718293M33QHDEQBROL"
	dnsRecordsFile     = "assets/samples/records.json"
	dnsRecordsYamlFile = "assets/samples/records.txt"
	filterTXT          = "TXT"
)

func TestClassifierFailWrongKind(t *testing.T) {
	g := NewWithT(t)
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(RightRawData)
	if err != nil {
		panic("Error encoding data")
	}
	err = Classifier(RightFormat, &buff, RightPath, WrongKind)
	g.Expect(err).Should(MatchError("Kind type not implemented: \n - rev"))
}

func TestClassifierFailDecodingRec(t *testing.T) {
	g := NewWithT(t)
	var buff bytes.Buffer
	xChanges := make([]Changes, 0, 0)
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", "SampleComment", RightZoneID),
		Changes: xChanges,
	})
	if err != nil {
		panic("Error encoding data")
	}
	err = Classifier(RightFormat, &buff, RightPath, RecRightKind)
	g.Expect(err).Should(MatchError("Error decoding data from buffer: \n - gob: decoding into local type *awsRoute53BG.localroute53RRS, received remote type ChangeJson = struct { Comment string; Changes []Changes = struct { Action string; ResourceRecordSets []CResourceRecordSets = struct { Name string; Type string; }; }; }"))
}

func TestClassifierFailDecodingWrongPath(t *testing.T) {
	g := NewWithT(t)
	var buff bytes.Buffer
	xChanges := make([]Changes, 0, 0)
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", "SampleComment", RightZoneID),
		Changes: xChanges,
	})
	if err != nil {
		panic("Error encoding data")
	}
	err = Classifier(RightFormat, &buff, WrongPath, GenRightKind)
	g.Expect(err).Should(MatchError("Error exporting data to struct: \n - Error exporting to JSON: \n - Error writing data into file /temepe/fail.json: \n - open /temepe/fail.json: no such file or directory"))
}

func TestClassifierSuccessDecodingGen(t *testing.T) {
	g := NewWithT(t)
	var buff bytes.Buffer
	xChanges := make([]Changes, 0, 0)
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", "SampleComment", RightZoneID),
		Changes: xChanges,
	})
	if err != nil {
		panic("Error encoding data")
	}
	err = Classifier(RightFormat, &buff, RightPath, GenRightKind)
	g.Expect(err).Should(BeNil())
}

func TestExporterSuccess(t *testing.T) {
	g := NewWithT(t)
	var buff bytes.Buffer
	var RRS localroute53RRS
	xChanges := make([]Changes, 0, 0)
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", "SampleComment", RightZoneID),
		Changes: xChanges,
	})
	if err != nil {
		panic("Error encoding data")
	}
	dec := gob.NewDecoder(&buff)

	err = dec.Decode(&RRS)
	err = Exporter(RightFormat, RightPath, RecRightKind, RRS)
	g.Expect(err).Should(BeNil())
}

func TestExporterFailWrongFormat(t *testing.T) {
	g := NewWithT(t)
	var buff bytes.Buffer
	var RRS localroute53RRS
	xChanges := make([]Changes, 0, 0)
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", "SampleComment", RightZoneID),
		Changes: xChanges,
	})
	if err != nil {
		panic("Error encoding data")
	}
	dec := gob.NewDecoder(&buff)

	err = dec.Decode(&RRS)
	err = Exporter(WrongFormat, RightPath, RecRightKind, &RRS)
	g.Expect(err).ShouldNot(BeNil())
}

func TestExporterFailWrongPath(t *testing.T) {
	g := NewWithT(t)
	var buff bytes.Buffer
	var RRS localroute53RRS
	xChanges := make([]Changes, 0, 0)
	enc := gob.NewEncoder(&buff)

	err := enc.Encode(ChangeJson{
		Comment: fmt.Sprintf("%s in Zone: %v", "SampleComment", RightZoneID),
		Changes: xChanges,
	})
	if err != nil {
		panic("Error encoding data")
	}
	dec := gob.NewDecoder(&buff)

	err = dec.Decode(&RRS)
	err = Exporter(RightFormat, WrongPath, RecRightKind, &RRS)
	g.Expect(err).ShouldNot(BeNil())
}
