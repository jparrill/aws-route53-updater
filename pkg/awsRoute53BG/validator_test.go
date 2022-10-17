package awsRoute53BG

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestValidateActionFail(t *testing.T) {
	g := NewWithT(t)
	err := ValidateAction(WrongAction)
	g.Expect(err).Should(MatchError("Action not supported by AWS Route53 service, only (CREATE|UPSERT|DELETE) are implemented. Called: DILATE\n"))
}

func TestValidateActionSuccess(t *testing.T) {
	g := NewWithT(t)
	err := ValidateAction(RightAction)
	g.Expect(err).Should(BeNil())
}

func TestValidateZoneIDFail(t *testing.T) {
	g := NewWithT(t)
	err := ValidateZoneID(WrongZoneID)
	g.Expect(err).ShouldNot(BeNil())
	g.Expect(err.Error()).To(ContainSubstring("Error"))
}

func TestValidateZoneIDSuccess(t *testing.T) {
	g := NewWithT(t)
	err := ValidateZoneID(RightZoneID)
	g.Expect(err).Should(BeNil())
}
