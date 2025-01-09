// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileStepParamsElement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#context SyntheticsTest#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Valid values are `native`, `web`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#context_type SyntheticsTest#context_type}
	ContextType *string `field:"optional" json:"contextType" yaml:"contextType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#element_description SyntheticsTest#element_description}.
	ElementDescription *string `field:"optional" json:"elementDescription" yaml:"elementDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#multi_locator SyntheticsTest#multi_locator}.
	MultiLocator *map[string]*string `field:"optional" json:"multiLocator" yaml:"multiLocator"`
	// relative_position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#relative_position SyntheticsTest#relative_position}
	RelativePosition *SyntheticsTestMobileStepParamsElementRelativePosition `field:"optional" json:"relativePosition" yaml:"relativePosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#text_content SyntheticsTest#text_content}.
	TextContent *string `field:"optional" json:"textContent" yaml:"textContent"`
	// user_locator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#user_locator SyntheticsTest#user_locator}
	UserLocator *SyntheticsTestMobileStepParamsElementUserLocator `field:"optional" json:"userLocator" yaml:"userLocator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/synthetics_test#view_name SyntheticsTest#view_name}.
	ViewName *string `field:"optional" json:"viewName" yaml:"viewName"`
}

