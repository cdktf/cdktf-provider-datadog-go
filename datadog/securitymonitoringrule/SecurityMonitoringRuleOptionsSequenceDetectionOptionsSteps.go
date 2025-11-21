// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsSequenceDetectionOptionsSteps struct {
	// Condition for the step to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/security_monitoring_rule#condition SecurityMonitoringRule#condition}
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Unique name of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Evaluation window for the step.
	//
	// Valid values are `0`, `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `10800`, `21600`, `43200`, `86400`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/security_monitoring_rule#evaluation_window SecurityMonitoringRule#evaluation_window}
	EvaluationWindow *float64 `field:"optional" json:"evaluationWindow" yaml:"evaluationWindow"`
}

