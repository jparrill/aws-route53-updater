package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/jparrill/aws-route53-updater/pkg/awsRoute53BG"
)

func RecoverHostedZone(zoneID, outputPath, outputFormat string, filters ...string) {
	svc := route53.New(session.New())
	input := &route53.GetHostedZoneInput{
		Id: aws.String(zoneID),
	}

	result, err := svc.GetHostedZone(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Println(route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Println(route53.ErrCodeInvalidInput, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)

}

func RecoverRecordSet(zoneID string) ([]*route53.ResourceRecordSet, error) {
	svc := route53.New(session.New())
	input := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	result, err := svc.ListResourceRecordSets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case route53.ErrCodeNoSuchHostedZone:
				fmt.Errorf("Error Zone does not exists: %v\n, trace: %v\n", route53.ErrCodeNoSuchHostedZone, aerr.Error())
			case route53.ErrCodeInvalidInput:
				fmt.Errorf("Error Code: %v\n, trace: %v\n", route53.ErrCodeInvalidInput, aerr.Error())
			default:
				return result.ResourceRecordSets, fmt.Errorf("Error not managed in data received from AWS request: %v\n", aerr.Error())
			}
		} else {
			return result.ResourceRecordSets, fmt.Errorf("Error in data received from AWS request: %v\n", err.Error())
		}
	}

	return result.ResourceRecordSets, nil
}

func Recover(zoneID, outputPath, outputFormat string, filters ...string) error {
	var buff bytes.Buffer
	kind := "rec"
	enc := gob.NewEncoder(&buff)
	rrs, err := RecoverRecordSet(zoneID)
	if err != nil {
		return fmt.Errorf("Error recovering record set: %v\n", err)
	}
	err = enc.Encode(rrs)
	if err != nil {
		return fmt.Errorf("Error storing data for processing into buffer: %v\n", err)
	}

	err = awsRoute53BG.Classifier(outputFormat, &buff, outputPath, kind)
	if err != nil {
		return fmt.Errorf("Error classifying or exporting data received: %v\n", err)
	}

	return nil
}
