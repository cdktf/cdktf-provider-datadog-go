package securitymonitoringrule


type SecurityMonitoringRuleFilter struct {
	// The type of filtering action. Valid values are `require`, `suppress`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#action SecurityMonitoringRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Query for selecting logs to apply the filtering action.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#query SecurityMonitoringRule#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

