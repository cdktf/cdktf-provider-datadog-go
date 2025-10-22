// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileStepParams struct {
	// Check type to use for an assertion step.
	//
	// Valid values are `equals`, `notEquals`, `contains`, `notContains`, `startsWith`, `notStartsWith`, `greater`, `lower`, `greaterEquals`, `lowerEquals`, `matchRegex`, `between`, `isEmpty`, `notIsEmpty`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#check SyntheticsTest#check}
	Check *string `field:"optional" json:"check" yaml:"check"`
	// Delay between each key stroke for a "type test" step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#delay SyntheticsTest#delay}
	Delay *float64 `field:"optional" json:"delay" yaml:"delay"`
	// Valid values are `up`, `down`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#direction SyntheticsTest#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// element block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#element SyntheticsTest#element}
	Element *SyntheticsTestMobileStepParamsElement `field:"optional" json:"element" yaml:"element"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#enable SyntheticsTest#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#max_scrolls SyntheticsTest#max_scrolls}.
	MaxScrolls *float64 `field:"optional" json:"maxScrolls" yaml:"maxScrolls"`
	// positions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#positions SyntheticsTest#positions}
	Positions interface{} `field:"optional" json:"positions" yaml:"positions"`
	// ID of the Synthetics test to use as subtest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#subtest_public_id SyntheticsTest#subtest_public_id}
	SubtestPublicId *string `field:"optional" json:"subtestPublicId" yaml:"subtestPublicId"`
	// Value of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#value SyntheticsTest#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#variable SyntheticsTest#variable}
	Variable *SyntheticsTestMobileStepParamsVariable `field:"optional" json:"variable" yaml:"variable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#with_enter SyntheticsTest#with_enter}.
	WithEnter interface{} `field:"optional" json:"withEnter" yaml:"withEnter"`
	// X coordinates for a "scroll step".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#x SyntheticsTest#x}
	X *float64 `field:"optional" json:"x" yaml:"x"`
	// Y coordinates for a "scroll step".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/synthetics_test#y SyntheticsTest#y}
	Y *float64 `field:"optional" json:"y" yaml:"y"`
}

