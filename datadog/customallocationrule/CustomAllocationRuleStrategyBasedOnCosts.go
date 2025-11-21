// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrule


type CustomAllocationRuleStrategyBasedOnCosts struct {
	// The condition to match. Valid values are `=`, `!=`, `is`, `is not`, `like`, `in`, `not in`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/custom_allocation_rule#condition CustomAllocationRule#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The tag key to use as the basis for cost allocation calculations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/custom_allocation_rule#tag CustomAllocationRule#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The single tag value to use for cost calculations. Use with conditions like `=`, `!=`, `is`, `is not`, `like`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/custom_allocation_rule#value CustomAllocationRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// A list of tag values to use for cost calculations. Use with `in` or `not in` conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/custom_allocation_rule#values CustomAllocationRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

