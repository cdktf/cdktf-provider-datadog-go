// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestApiStepAssertionTargetjsonpath struct {
	// The JSON path to assert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_test#jsonpath SyntheticsTest#jsonpath}
	Jsonpath *string `field:"required" json:"jsonpath" yaml:"jsonpath"`
	// The specific operator to use on the path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_test#operator SyntheticsTest#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The element from the list of results to assert on.
	//
	// Select from `firstElementMatches` (the first element in the list), `everyElementMatches` (every element in the list), `atLeastOneElementMatches` (at least one element in the list), or `serializationMatches` (the serialized value of the list). Defaults to `"firstElementMatches"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_test#elementsoperator SyntheticsTest#elementsoperator}
	Elementsoperator *string `field:"optional" json:"elementsoperator" yaml:"elementsoperator"`
	// Expected matching value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_test#targetvalue SyntheticsTest#targetvalue}
	Targetvalue *string `field:"optional" json:"targetvalue" yaml:"targetvalue"`
}

