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

func RecoverRecordSet(zoneID string) []*route53.ResourceRecordSet {
	svc := route53.New(session.New())
	input := &route53.ListResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneID),
	}

	result, err := svc.ListResourceRecordSets(input)
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
	}

	return result.ResourceRecordSets
}

func Recover(zoneID, outputPath, outputFormat string, filters ...string) {
	var buff bytes.Buffer
	kind := "rec"
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(RecoverRecordSet(zoneID))
	if err != nil {
		panic(err)
	}

	awsRoute53BG.Classifier(outputFormat, &buff, outputPath, kind)
}
