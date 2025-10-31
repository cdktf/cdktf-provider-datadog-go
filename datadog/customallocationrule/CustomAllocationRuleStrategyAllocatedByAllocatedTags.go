// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrule


type CustomAllocationRuleStrategyAllocatedByAllocatedTags struct {
	// The tag key to allocate costs to (e.g., `team`, `environment`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/custom_allocation_rule#key CustomAllocationRule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value to allocate costs to (e.g., `backend`, `production`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/custom_allocation_rule#value CustomAllocationRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

