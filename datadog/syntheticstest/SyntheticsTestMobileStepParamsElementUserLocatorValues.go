// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileStepParamsElementUserLocatorValues struct {
	// Valid values are `accessibility-id`, `id`, `ios-predicate-string`, `ios-class-chain`, `xpath`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#value SyntheticsTest#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

