package internal

import (
	"testing"

	. "github.com/onsi/gomega"
)

const (
	RecWrongZoneID       = "SICK"
	RecWrongOutputPath   = "/temepe/fail.json"
	RecWrongOutputFormat = "jason"
	RecRightZoneID       = "Z02718293M33QHDEQBROL"
	RecRightOutputPath   = "/tmp/test.json"
	RecRightOutputFormat = "json"
)

func TestRecoverFailWrongZone(t *testing.T) {
	g := NewWithT(t)
	err := Recover(RecWrongZoneID, RecRightOutputPath, RecRightOutputFormat)
	g.Expect(err).Should(HaveOccurred(), "Error recovering record set")
}

func TestRecoverFailWrongOutFormat(t *testing.T) {
	g := NewWithT(t)
	err := Recover(RecRightZoneID, RecRightOutputPath, RecWrongOutputFormat)
	g.Expect(err).Should(HaveOccurred(), "Output format not implemented")
}

func TestRecoverFailWrongOutPath(t *testing.T) {
	g := NewWithT(t)
	err := Recover(RecRightZoneID, RecWrongOutputPath, RecRightOutputFormat)
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(err).Should(HaveOccurred(), "Error classifying or exporting data received")
}

func TestRecoverRight(t *testing.T) {
	g := NewWithT(t)
	err := Recover(RecRightZoneID, RecRightOutputPath, RecRightOutputFormat)
	g.Expect(err).Should(BeNil())
}

func TestRecoverRecordSetRecRightZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(RecRightZoneID)
	g.Expect(err).Should(BeNil())
	g.Expect(rrs).ShouldNot(BeEmpty())
}

func TestRecoverRecordSetRecWrongZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(RecWrongZoneID)
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(rrs).Should(BeEmpty())
}

func TestRecoverHostedZoneRecRightZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(RecRightZoneID)
	g.Expect(err).Should(BeNil())
	g.Expect(rrs).ShouldNot(BeEmpty())
}

func TestRecoverHostedZoneRecWrongZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(RecWrongZoneID)
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(rrs).Should(BeEmpty())
}
