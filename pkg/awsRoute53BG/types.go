package awsRoute53BG

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
