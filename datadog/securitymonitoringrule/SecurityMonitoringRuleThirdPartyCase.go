// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleThirdPartyCase struct {
	// Severity of the Security Signal. Valid values are `info`, `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/security_monitoring_rule#status SecurityMonitoringRule#status}
	Status *string `field:"required" json:"status" yaml:"status"`
	// Name of the case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Notification targets for each rule case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/security_monitoring_rule#notifications SecurityMonitoringRule#notifications}
	Notifications *[]*string `field:"optional" json:"notifications" yaml:"notifications"`
	// A query to associate a third-party event to this case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/security_monitoring_rule#query SecurityMonitoringRule#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

