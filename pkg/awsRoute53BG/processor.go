package awsRoute53BG

import "golang.org/x/exp/slices"

func ProcessData(action string, data []ResourceRecordSets, filters ...string) Changes {
	newData := make([]CResourceRecordSets, 0, 0)

	if len(filters) == 0 {
		for _, v := range data {
			if v.Type != "SOA" && v.Type != "NS" {
				newEntry := CResourceRecordSets{
					Name: v.Name,
					Type: v.Type,
				}
				newData = append(newData, newEntry)

			}
		}
	} else {
		for _, v := range data {
			if v.Type != "SOA" && v.Type != "NS" {
				if slices.Contains(filters, v.Type) {
					newEntry := CResourceRecordSets{
						Name: v.Name,
						Type: v.Type,
					}
					newData = append(newData, newEntry)
				}
			}
		}
	}

	c := Changes{
		Action:             action,
		ResourceRecordSets: newData,
	}

	return c
}
