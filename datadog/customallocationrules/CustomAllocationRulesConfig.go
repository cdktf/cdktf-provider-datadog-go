// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomAllocationRulesConfig struct {
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
	// The list of Custom Allocation Rule IDs, in order.
	//
	// Rules are executed in the order specified in this list. Comes from the `id` field on a `datadog_custom_allocation_rule` resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/custom_allocation_rules#rule_ids CustomAllocationRules#rule_ids}
	RuleIds *[]*string `field:"required" json:"ruleIds" yaml:"ruleIds"`
}

