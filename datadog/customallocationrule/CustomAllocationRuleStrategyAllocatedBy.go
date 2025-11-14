// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrule


type CustomAllocationRuleStrategyAllocatedBy struct {
	// allocated_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/custom_allocation_rule#allocated_tags CustomAllocationRule#allocated_tags}
	AllocatedTags interface{} `field:"optional" json:"allocatedTags" yaml:"allocatedTags"`
	// The percentage of costs to allocate to this target as a decimal (e.g., 0.33 for 33%). Used when `method` is `percent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/custom_allocation_rule#percentage CustomAllocationRule#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

