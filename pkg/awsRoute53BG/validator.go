package awsRoute53BG

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

func ValidateAction(action string) error {
	switch action {
	case "UPSERT", "DELETE", "CREATE":
		return nil
	default:
		return fmt.Errorf("Action not supported by AWS Route53 service, only (CREATE|UPSERT|DELETE) are implemented. Called: %s\n", action)
	}
}

func ValidateZoneID(zoneID string) error {
	svc := route53.New(session.New())
	input := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	_, err := svc.ListResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				return fmt.Errorf("Error, Zone does not exists: %v\n, trace: %v\n", route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeInvalidInput:
				return fmt.Errorf("Error Code: %v\n, trace: %v\n", route53.ErrCodeInvalidInput, aerr.Error())
			default:
				return fmt.Errorf("Error not managed in data received from AWS request: \n - %v", aerr.Error())
			}
		}
	}
	return nil
}
