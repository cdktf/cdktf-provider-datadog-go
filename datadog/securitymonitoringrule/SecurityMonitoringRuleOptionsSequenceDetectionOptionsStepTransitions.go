// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsSequenceDetectionOptionsStepTransitions struct {
	// Child step name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/security_monitoring_rule#child SecurityMonitoringRule#child}
	Child *string `field:"required" json:"child" yaml:"child"`
	// Parent step name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/security_monitoring_rule#parent SecurityMonitoringRule#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Maximum time allowed to transition from parent to child.
	//
	// Valid values are `0`, `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `10800`, `21600`, `43200`, `86400`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/security_monitoring_rule#evaluation_window SecurityMonitoringRule#evaluation_window}
	EvaluationWindow *float64 `field:"optional" json:"evaluationWindow" yaml:"evaluationWindow"`
}

