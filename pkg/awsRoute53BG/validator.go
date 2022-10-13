package awsRoute53BG

import "fmt"

func ValidateAction(action string) error {
	switch action {
	case "UPSERT", "DELETE", "CREATE":
		return nil
	default:
		return fmt.Errorf("Action not supported by AWS Route53 service, only (CREATE|UPSERT|DELETE) are implemented. Called: %s\n", action)
	}
}
