// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogcustomallocationrule


type DataDatadogCustomAllocationRuleStrategy struct {
	// allocated_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/data-sources/custom_allocation_rule#allocated_by DataDatadogCustomAllocationRule#allocated_by}
	AllocatedBy interface{} `field:"optional" json:"allocatedBy" yaml:"allocatedBy"`
	// allocated_by_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/data-sources/custom_allocation_rule#allocated_by_filters DataDatadogCustomAllocationRule#allocated_by_filters}
	AllocatedByFilters interface{} `field:"optional" json:"allocatedByFilters" yaml:"allocatedByFilters"`
	// based_on_costs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/data-sources/custom_allocation_rule#based_on_costs DataDatadogCustomAllocationRule#based_on_costs}
	BasedOnCosts interface{} `field:"optional" json:"basedOnCosts" yaml:"basedOnCosts"`
	// based_on_timeseries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/data-sources/custom_allocation_rule#based_on_timeseries DataDatadogCustomAllocationRule#based_on_timeseries}
	BasedOnTimeseries *DataDatadogCustomAllocationRuleStrategyBasedOnTimeseries `field:"optional" json:"basedOnTimeseries" yaml:"basedOnTimeseries"`
	// evaluate_grouped_by_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/data-sources/custom_allocation_rule#evaluate_grouped_by_filters DataDatadogCustomAllocationRule#evaluate_grouped_by_filters}
	EvaluateGroupedByFilters interface{} `field:"optional" json:"evaluateGroupedByFilters" yaml:"evaluateGroupedByFilters"`
}

