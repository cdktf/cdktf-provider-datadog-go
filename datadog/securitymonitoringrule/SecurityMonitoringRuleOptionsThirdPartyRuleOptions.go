// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsThirdPartyRuleOptions struct {
	// Severity of the default rule case, when none of the third-party cases match.
	//
	// Valid values are `info`, `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/security_monitoring_rule#default_status SecurityMonitoringRule#default_status}
	DefaultStatus *string `field:"required" json:"defaultStatus" yaml:"defaultStatus"`
	// root_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/security_monitoring_rule#root_query SecurityMonitoringRule#root_query}
	RootQuery interface{} `field:"required" json:"rootQuery" yaml:"rootQuery"`
	// Notification targets for the default rule case, when none of the third-party cases match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/security_monitoring_rule#default_notifications SecurityMonitoringRule#default_notifications}
	DefaultNotifications *[]*string `field:"optional" json:"defaultNotifications" yaml:"defaultNotifications"`
	// A template for the signal title; if omitted, the title is generated based on the case name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/security_monitoring_rule#signal_title_template SecurityMonitoringRule#signal_title_template}
	SignalTitleTemplate *string `field:"optional" json:"signalTitleTemplate" yaml:"signalTitleTemplate"`
}

