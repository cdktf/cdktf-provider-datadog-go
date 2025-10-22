// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringdefaultrule


type SecurityMonitoringDefaultRuleCase struct {
	// Status of the rule case to match. Valid values are `info`, `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/security_monitoring_default_rule#status SecurityMonitoringDefaultRule#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Status of the rule case to override. Valid values are `info`, `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/security_monitoring_default_rule#custom_status SecurityMonitoringDefaultRule#custom_status}
	CustomStatus *string `field:"optional" json:"customStatus" yaml:"customStatus"`
	// Notification targets for each rule case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/security_monitoring_default_rule#notifications SecurityMonitoringDefaultRule#notifications}
	Notifications *[]*string `field:"optional" json:"notifications" yaml:"notifications"`
}

