// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customallocationrule


type CustomAllocationRuleCostsToAllocate struct {
	// The condition to match. Valid values are `=`, `!=`, `is`, `is not`, `like`, `in`, `not in`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/custom_allocation_rule#condition CustomAllocationRule#condition}
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// The tag key to filter on (e.g., `aws_product`, `team`, `environment`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/custom_allocation_rule#tag CustomAllocationRule#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The single tag value to match.
	//
	// Use this field for conditions like `=`, `!=`, `is`, `is not`, `like`. Do not use with `in` or `not in` conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/custom_allocation_rule#value CustomAllocationRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// A list of tag values to match.
	//
	// Use this field for `in` or `not in` conditions only. Do not use with single-value conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/custom_allocation_rule#values CustomAllocationRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

