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
	// Indicates Whether the Agent rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/csm_threats_agent_rule#enabled CsmThreatsAgentRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The SECL expression of the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/csm_threats_agent_rule#expression CsmThreatsAgentRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The name of the Agent rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/csm_threats_agent_rule#name CsmThreatsAgentRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the Agent rule. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/csm_threats_agent_rule#description CsmThreatsAgentRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

