// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csmthreatsagentrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CsmThreatsAgentRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The SECL expression of the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_agent_rule#expression CsmThreatsAgentRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_agent_rule#name CsmThreatsAgentRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_agent_rule#actions CsmThreatsAgentRule#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// A description for the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_agent_rule#description CsmThreatsAgentRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the Agent rule is enabled. Must not be used without policy_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_agent_rule#enabled CsmThreatsAgentRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID of the agent policy in which the rule is saved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_agent_rule#policy_id CsmThreatsAgentRule#policy_id}
	PolicyId *string `field:"optional" json:"policyId" yaml:"policyId"`
	// The list of product tags associated with the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_agent_rule#product_tags CsmThreatsAgentRule#product_tags}
	ProductTags *[]*string `field:"optional" json:"productTags" yaml:"productTags"`
}

