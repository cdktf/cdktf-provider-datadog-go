// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallescalationpolicy


type OnCallEscalationPolicyStep struct {
	// Defines how many seconds to wait before escalating to the next step. Value must be between 60 and 36000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/on_call_escalation_policy#escalate_after_seconds OnCallEscalationPolicy#escalate_after_seconds}
	EscalateAfterSeconds *float64 `field:"required" json:"escalateAfterSeconds" yaml:"escalateAfterSeconds"`
	// Specifies how this escalation step will assign targets.
	//
	// Can be `default` (page all targets at once) or `round-robin`. Valid values are `default`, `round-robin`. Defaults to `"default"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/on_call_escalation_policy#assignment OnCallEscalationPolicy#assignment}
	Assignment *string `field:"optional" json:"assignment" yaml:"assignment"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/on_call_escalation_policy#target OnCallEscalationPolicy#target}
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

