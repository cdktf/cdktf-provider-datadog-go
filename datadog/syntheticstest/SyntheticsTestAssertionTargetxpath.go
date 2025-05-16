// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestAssertionTargetxpath struct {
	// The specific operator to use on the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/synthetics_test#operator SyntheticsTest#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The xpath to assert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/synthetics_test#xpath SyntheticsTest#xpath}
	Xpath *string `field:"required" json:"xpath" yaml:"xpath"`
	// Expected matching value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/synthetics_test#targetvalue SyntheticsTest#targetvalue}
	Targetvalue *string `field:"optional" json:"targetvalue" yaml:"targetvalue"`
}

