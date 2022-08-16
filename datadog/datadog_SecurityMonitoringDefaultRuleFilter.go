// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SecurityMonitoringDefaultRuleFilter struct {
	// The type of filtering action. Allowed enum values: require, suppress Valid values are `require`, `suppress`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_default_rule#action SecurityMonitoringDefaultRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Query for selecting logs to apply the filtering action.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_default_rule#query SecurityMonitoringDefaultRule#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

