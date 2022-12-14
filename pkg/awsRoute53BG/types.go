package awsRoute53BG

import "github.com/aws/aws-sdk-go/service/route53"

type JSONFile struct {
	Format   string
	FilePath string
	Data     []byte
}

type YAMLFile struct {
	Format   string
	FilePath string
	Data     []byte
}

type ChangeJson struct {
	Comment string    `json:"Comment"`
	Changes []Changes `json:"Changes"`
}

type Changes struct {
	Action             string                `json:"Action"`
	ResourceRecordSets []CResourceRecordSets `json:"ResourceRecordSets"`
}
type AWSRecords struct {
	ResourceRecordSets []ResourceRecordSets `json:"ResourceRecordSets"`
}
type ResourceRecords struct {
	Value string `json:"Value"`
}
type CResourceRecordSets struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
}
type ResourceRecordSets struct {
	Name            string            `json:"Name"`
	Type            string            `json:"Type"`
	TTL             int               `json:"TTL"`
	ResourceRecords []ResourceRecords `json:"ResourceRecords"`
}
type localroute53RRS []*route53.ResourceRecordSet
