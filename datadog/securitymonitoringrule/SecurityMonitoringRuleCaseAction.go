// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleCaseAction struct {
	// Type of action to perform when the case triggers. Valid values are `block_ip`, `block_user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/security_monitoring_rule#type SecurityMonitoringRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/security_monitoring_rule#options SecurityMonitoringRule#options}
	Options *SecurityMonitoringRuleCaseActionOptions `field:"optional" json:"options" yaml:"options"`
}

