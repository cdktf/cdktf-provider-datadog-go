// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsSequenceDetectionOptions struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/security_monitoring_rule#steps SecurityMonitoringRule#steps}
	Steps interface{} `field:"optional" json:"steps" yaml:"steps"`
	// step_transitions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/security_monitoring_rule#step_transitions SecurityMonitoringRule#step_transitions}
	StepTransitions interface{} `field:"optional" json:"stepTransitions" yaml:"stepTransitions"`
}

