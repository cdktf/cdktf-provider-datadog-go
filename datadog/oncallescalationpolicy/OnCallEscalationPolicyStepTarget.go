// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallescalationpolicy


type OnCallEscalationPolicyStepTarget struct {
	// Targeted schedule ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/on_call_escalation_policy#schedule OnCallEscalationPolicy#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Targeted team ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/on_call_escalation_policy#team OnCallEscalationPolicy#team}
	Team *string `field:"optional" json:"team" yaml:"team"`
	// Targeted user ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/on_call_escalation_policy#user OnCallEscalationPolicy#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

