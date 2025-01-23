// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestBrowserStep struct {
	// Name of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#params SyntheticsTest#params}
	Params *SyntheticsTestBrowserStepParams `field:"required" json:"params" yaml:"params"`
	// Type of the step.
	//
	// Valid values are `assertCurrentUrl`, `assertElementAttribute`, `assertElementContent`, `assertElementPresent`, `assertEmail`, `assertFileDownload`, `assertFromJavascript`, `assertPageContains`, `assertPageLacks`, `click`, `extractFromJavascript`, `extractVariable`, `goToEmailLink`, `goToUrl`, `goToUrlAndMeasureTti`, `hover`, `playSubTest`, `pressKey`, `refresh`, `runApiTest`, `scroll`, `selectOption`, `typeText`, `uploadFiles`, `wait`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Determines if the step should be allowed to fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#allow_failure SyntheticsTest#allow_failure}
	AllowFailure interface{} `field:"optional" json:"allowFailure" yaml:"allowFailure"`
	// Determines whether or not to always execute this step even if the previous step failed or was skipped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#always_execute SyntheticsTest#always_execute}
	AlwaysExecute interface{} `field:"optional" json:"alwaysExecute" yaml:"alwaysExecute"`
	// Determines whether or not to exit the test if the step succeeds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#exit_if_succeed SyntheticsTest#exit_if_succeed}
	ExitIfSucceed interface{} `field:"optional" json:"exitIfSucceed" yaml:"exitIfSucceed"`
	// Force update of the "element" parameter for the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#force_element_update SyntheticsTest#force_element_update}
	ForceElementUpdate interface{} `field:"optional" json:"forceElementUpdate" yaml:"forceElementUpdate"`
	// Determines whether or not to consider the entire test as failed if this step fails.
	//
	// Can be used only if `allow_failure` is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#is_critical SyntheticsTest#is_critical}
	IsCritical interface{} `field:"optional" json:"isCritical" yaml:"isCritical"`
	// A unique identifier used to track steps after reordering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#local_key SyntheticsTest#local_key}
	LocalKey *string `field:"optional" json:"localKey" yaml:"localKey"`
	// Prevents saving screenshots of the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#no_screenshot SyntheticsTest#no_screenshot}
	NoScreenshot interface{} `field:"optional" json:"noScreenshot" yaml:"noScreenshot"`
	// Used to override the default timeout of a step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.53.0/docs/resources/synthetics_test#timeout SyntheticsTest#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

