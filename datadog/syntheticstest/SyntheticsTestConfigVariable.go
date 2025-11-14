// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestConfigVariable struct {
	// Name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of test configuration variable. Valid values are `global`, `text`, `email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Example for the variable.
	//
	// This value is not returned by the API when `secure = true`. Avoid drift by only making updates to this value from within Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#example SyntheticsTest#example}
	Example *string `field:"optional" json:"example" yaml:"example"`
	// When type = `global`, ID of the global variable to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#id SyntheticsTest#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Pattern of the variable.
	//
	// This value is not returned by the API when `secure = true`. Avoid drift by only making updates to this value from within Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#pattern SyntheticsTest#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Whether the value of this variable will be obfuscated in test results. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/synthetics_test#secure SyntheticsTest#secure}
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
}

