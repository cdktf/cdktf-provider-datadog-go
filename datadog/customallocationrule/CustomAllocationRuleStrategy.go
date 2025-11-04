// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrule


type CustomAllocationRuleStrategy struct {
	// allocated_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#allocated_by CustomAllocationRule#allocated_by}
	AllocatedBy interface{} `field:"optional" json:"allocatedBy" yaml:"allocatedBy"`
	// allocated_by_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#allocated_by_filters CustomAllocationRule#allocated_by_filters}
	AllocatedByFilters interface{} `field:"optional" json:"allocatedByFilters" yaml:"allocatedByFilters"`
	// List of tag keys used to allocate costs (e.g., `["team", "project"]`). Costs will be distributed across unique values of these tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#allocated_by_tag_keys CustomAllocationRule#allocated_by_tag_keys}
	AllocatedByTagKeys *[]*string `field:"optional" json:"allocatedByTagKeys" yaml:"allocatedByTagKeys"`
	// based_on_costs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#based_on_costs CustomAllocationRule#based_on_costs}
	BasedOnCosts interface{} `field:"optional" json:"basedOnCosts" yaml:"basedOnCosts"`
	// based_on_timeseries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#based_on_timeseries CustomAllocationRule#based_on_timeseries}
	BasedOnTimeseries *CustomAllocationRuleStrategyBasedOnTimeseries `field:"optional" json:"basedOnTimeseries" yaml:"basedOnTimeseries"`
	// evaluate_grouped_by_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#evaluate_grouped_by_filters CustomAllocationRule#evaluate_grouped_by_filters}
	EvaluateGroupedByFilters interface{} `field:"optional" json:"evaluateGroupedByFilters" yaml:"evaluateGroupedByFilters"`
	// List of tag keys used to group costs before allocation.
	//
	// Costs are grouped by these tag values before applying the allocation strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#evaluate_grouped_by_tag_keys CustomAllocationRule#evaluate_grouped_by_tag_keys}
	EvaluateGroupedByTagKeys *[]*string `field:"optional" json:"evaluateGroupedByTagKeys" yaml:"evaluateGroupedByTagKeys"`
	// The granularity level for cost allocation. Valid values are `daily` or `monthly`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#granularity CustomAllocationRule#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// The allocation method. Valid values are `even`, `proportional`, `proportional_timeseries`, or `percent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/custom_allocation_rule#method CustomAllocationRule#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
}

