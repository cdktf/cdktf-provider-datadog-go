// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleFilter struct {
	// The type of filtering action. Valid values are `require`, `suppress`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/security_monitoring_rule#action SecurityMonitoringRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Query for selecting logs to apply the filtering action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/security_monitoring_rule#query SecurityMonitoringRule#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

