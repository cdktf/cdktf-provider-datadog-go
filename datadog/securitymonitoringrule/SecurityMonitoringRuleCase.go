package securitymonitoringrule


type SecurityMonitoringRuleCase struct {
	// Severity of the Security Signal. Valid values are `info`, `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#status SecurityMonitoringRule#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// A rule case contains logical operations (`>`,`>=`, `&&`, `||`) to determine if a signal should be generated based on the event counts in the previously defined queries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#condition SecurityMonitoringRule#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Name of the case.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Notification targets for each rule case.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#notifications SecurityMonitoringRule#notifications}
	Notifications *[]*string `field:"optional" json:"notifications" yaml:"notifications"`
}

