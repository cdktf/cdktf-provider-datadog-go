// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileStepParamsVariable struct {
	// Name of the extracted variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Example of the extracted variable. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/synthetics_test#example SyntheticsTest#example}
	Example *string `field:"optional" json:"example" yaml:"example"`
}

