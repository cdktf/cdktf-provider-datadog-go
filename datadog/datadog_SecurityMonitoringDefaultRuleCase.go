// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SecurityMonitoringDefaultRuleCase struct {
	// Notification targets for each rule case.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_default_rule#notifications SecurityMonitoringDefaultRule#notifications}
	Notifications *[]*string `field:"required" json:"notifications" yaml:"notifications"`
	// Status of the rule case to match. Valid values are `info`, `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_default_rule#status SecurityMonitoringDefaultRule#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

