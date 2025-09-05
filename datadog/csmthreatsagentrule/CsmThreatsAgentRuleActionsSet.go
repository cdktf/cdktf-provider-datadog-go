// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csmthreatsagentrule


type CsmThreatsAgentRuleActionsSet struct {
	// The name of the set action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/csm_threats_agent_rule#name CsmThreatsAgentRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to append to the set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/csm_threats_agent_rule#append CsmThreatsAgentRule#append}
	Append interface{} `field:"optional" json:"append" yaml:"append"`
	// The field to get the value from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/csm_threats_agent_rule#field CsmThreatsAgentRule#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The scope of the set action (process, container, cgroup, or empty).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/csm_threats_agent_rule#scope CsmThreatsAgentRule#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The maximum size of the set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/csm_threats_agent_rule#size CsmThreatsAgentRule#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The time to live for the set in nanoseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/csm_threats_agent_rule#ttl CsmThreatsAgentRule#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// The value to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/csm_threats_agent_rule#value CsmThreatsAgentRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

