// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrule


type CustomAllocationRuleStrategyEvaluateGroupedByFilters struct {
	// The condition to match. Valid values are `=`, `!=`, `is`, `is not`, `like`, `in`, `not in`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/custom_allocation_rule#condition CustomAllocationRule#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The tag key to filter on when grouping costs for evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/custom_allocation_rule#tag CustomAllocationRule#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The single tag value to match when grouping. Use with conditions like `=`, `!=`, `is`, `is not`, `like`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/custom_allocation_rule#value CustomAllocationRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// A list of tag values to match when grouping. Use with `in` or `not in` conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/custom_allocation_rule#values CustomAllocationRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

