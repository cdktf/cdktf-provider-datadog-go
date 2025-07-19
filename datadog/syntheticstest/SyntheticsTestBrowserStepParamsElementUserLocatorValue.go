// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestBrowserStepParamsElementUserLocatorValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/synthetics_test#value SyntheticsTest#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Defaults to `"css"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

