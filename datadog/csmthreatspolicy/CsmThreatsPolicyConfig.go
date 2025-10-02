// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csmthreatspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CsmThreatsPolicyConfig struct {
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
	// The name of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_policy#name CsmThreatsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_policy#description CsmThreatsPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether the policy is enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_policy#enabled CsmThreatsPolicy#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Host tags that define where the policy is deployed. Inner values are ANDed, outer arrays are ORed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_policy#host_tags_lists CsmThreatsPolicy#host_tags_lists}
	HostTagsLists interface{} `field:"optional" json:"hostTagsLists" yaml:"hostTagsLists"`
	// Host tags that define where the policy is deployed. Deprecated, use host_tags_lists instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/csm_threats_policy#tags CsmThreatsPolicy#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

