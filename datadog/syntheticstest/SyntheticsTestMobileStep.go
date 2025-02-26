// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileStep struct {
	// The name of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#params SyntheticsTest#params}
	Params *SyntheticsTestMobileStepParams `field:"required" json:"params" yaml:"params"`
	// The type of the step.
	//
	// Valid values are `assertElementContent`, `assertScreenContains`, `assertScreenLacks`, `doubleTap`, `extractVariable`, `flick`, `openDeeplink`, `playSubTest`, `pressBack`, `restartApplication`, `rotate`, `scroll`, `scrollToElement`, `tap`, `toggleWiFi`, `typeText`, `wait`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A boolean set to allow this step to fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#allow_failure SyntheticsTest#allow_failure}
	AllowFailure interface{} `field:"optional" json:"allowFailure" yaml:"allowFailure"`
	// A boolean set to determine if the step has a new step element.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#has_new_step_element SyntheticsTest#has_new_step_element}
	HasNewStepElement interface{} `field:"optional" json:"hasNewStepElement" yaml:"hasNewStepElement"`
	// A boolean to use in addition to `allowFailure` to determine if the test should be marked as failed when the step fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#is_critical SyntheticsTest#is_critical}
	IsCritical interface{} `field:"optional" json:"isCritical" yaml:"isCritical"`
	// A boolean set to not take a screenshot for the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#no_screenshot SyntheticsTest#no_screenshot}
	NoScreenshot interface{} `field:"optional" json:"noScreenshot" yaml:"noScreenshot"`
	// The public ID of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#public_id SyntheticsTest#public_id}
	PublicId *string `field:"optional" json:"publicId" yaml:"publicId"`
	// The time before declaring a step failed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#timeout SyntheticsTest#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

