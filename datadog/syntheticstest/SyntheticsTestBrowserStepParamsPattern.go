// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestBrowserStepParamsPattern struct {
	// Type of pattern to use for the step. Valid values are `regex`, `x_path`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Pattern to use for the step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/synthetics_test#value SyntheticsTest#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

