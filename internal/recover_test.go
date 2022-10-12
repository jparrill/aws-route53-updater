package internal

import (
	"testing"

	. "github.com/onsi/gomega"
)

const (
	WrongZoneID       = "SICK"
	WrongOutputPath   = "/temepe/fail.json"
	WrongOutputFormat = "jason"
	RightZoneID       = "Z02718293M33QHDEQBROL"
	RightOutputPath   = "/tmp/test.json"
	RightOutputFormat = "json"
)

func TestRecoverFailZoneWrong(t *testing.T) {
	g := NewWithT(t)
	err := Recover(WrongZoneID, RightOutputPath, RightOutputFormat)
	g.Expect(err).Should(HaveOccurred(), "Error recovering record set")
}

func TestRecoverFailOutFormatWrong(t *testing.T) {
	g := NewWithT(t)
	err := Recover(RightZoneID, RightOutputPath, WrongOutputFormat)
	g.Expect(err).Should(HaveOccurred(), "Output format not implemented")
}

func TestRecoverFailOutPathWrong(t *testing.T) {
	g := NewWithT(t)
	err := Recover(RightZoneID, WrongOutputPath, RightOutputFormat)
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(err).Should(HaveOccurred(), "Error classifying or exporting data received")
}

func TestRecoverRight(t *testing.T) {
	g := NewWithT(t)
	err := Recover(RightZoneID, RightOutputPath, RightOutputFormat)
	g.Expect(err).Should(BeNil())
}

func TestRecoverRecordSetRightZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(RightZoneID)
	g.Expect(err).Should(BeNil())
	g.Expect(rrs).ShouldNot(BeEmpty())
}

func TestRecoverRecordSetWrongZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(WrongZoneID)
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(rrs).Should(BeEmpty())
}

func TestRecoverHostedZoneRightZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(RightZoneID)
	g.Expect(err).Should(BeNil())
	g.Expect(rrs).ShouldNot(BeEmpty())
}

func TestRecoverHostedZoneWrongZoneID(t *testing.T) {
	g := NewWithT(t)
	rrs, err := RecoverRecordSet(WrongZoneID)
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(rrs).Should(BeEmpty())
}
