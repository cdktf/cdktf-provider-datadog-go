// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestBrowserStepParams struct {
	// Name of the attribute to use for an "assert attribute" step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#attribute SyntheticsTest#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Check type to use for an assertion step.
	//
	// Valid values are `equals`, `notEquals`, `contains`, `notContains`, `startsWith`, `notStartsWith`, `greater`, `lower`, `greaterEquals`, `lowerEquals`, `matchRegex`, `between`, `isEmpty`, `notIsEmpty`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#check SyntheticsTest#check}
	Check *string `field:"optional" json:"check" yaml:"check"`
	// Type of click to use for a "click" step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#click_type SyntheticsTest#click_type}
	ClickType *string `field:"optional" json:"clickType" yaml:"clickType"`
	// Javascript code to use for the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#code SyntheticsTest#code}
	Code *string `field:"optional" json:"code" yaml:"code"`
	// Delay between each key stroke for a "type test" step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#delay SyntheticsTest#delay}
	Delay *float64 `field:"optional" json:"delay" yaml:"delay"`
	// Element to use for the step, JSON encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#element SyntheticsTest#element}
	Element *string `field:"optional" json:"element" yaml:"element"`
	// element_user_locator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#element_user_locator SyntheticsTest#element_user_locator}
	ElementUserLocator *SyntheticsTestBrowserStepParamsElementUserLocator `field:"optional" json:"elementUserLocator" yaml:"elementUserLocator"`
	// Details of the email for an "assert email" step, JSON encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#email SyntheticsTest#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// JSON encoded string used for an "assert download" step.
	//
	// Refer to the examples for a usage example showing the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#file SyntheticsTest#file}
	File *string `field:"optional" json:"file" yaml:"file"`
	// Details of the files for an "upload files" step, JSON encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#files SyntheticsTest#files}
	Files *string `field:"optional" json:"files" yaml:"files"`
	// Modifier to use for a "press key" step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#modifiers SyntheticsTest#modifiers}
	Modifiers *[]*string `field:"optional" json:"modifiers" yaml:"modifiers"`
	// ID of the tab to play the subtest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#playing_tab_id SyntheticsTest#playing_tab_id}
	PlayingTabId *string `field:"optional" json:"playingTabId" yaml:"playingTabId"`
	// Request for an API step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#request SyntheticsTest#request}
	Request *string `field:"optional" json:"request" yaml:"request"`
	// ID of the Synthetics test to use as subtest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#subtest_public_id SyntheticsTest#subtest_public_id}
	SubtestPublicId *string `field:"optional" json:"subtestPublicId" yaml:"subtestPublicId"`
	// Value of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#value SyntheticsTest#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#variable SyntheticsTest#variable}
	Variable *SyntheticsTestBrowserStepParamsVariable `field:"optional" json:"variable" yaml:"variable"`
	// For "file upload" steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#with_click SyntheticsTest#with_click}
	WithClick interface{} `field:"optional" json:"withClick" yaml:"withClick"`
	// X coordinates for a "scroll step".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#x SyntheticsTest#x}
	X *float64 `field:"optional" json:"x" yaml:"x"`
	// Y coordinates for a "scroll step".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/synthetics_test#y SyntheticsTest#y}
	Y *float64 `field:"optional" json:"y" yaml:"y"`
}

