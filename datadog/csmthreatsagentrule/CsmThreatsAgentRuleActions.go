// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csmthreatsagentrule


type CsmThreatsAgentRuleActions struct {
	// set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/csm_threats_agent_rule#set CsmThreatsAgentRule#set}
	Set *CsmThreatsAgentRuleActionsSet `field:"optional" json:"set" yaml:"set"`
}

