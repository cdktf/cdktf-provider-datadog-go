// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestBrowserVariable struct {
	// Name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of browser test variable. Valid values are `element`, `email`, `global`, `text`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Example for the variable. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#example SyntheticsTest#example}
	Example *string `field:"optional" json:"example" yaml:"example"`
	// ID of the global variable to use.
	//
	// This is actually only used (and required) in the case of using a variable of type `global`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#id SyntheticsTest#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Pattern of the variable. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#pattern SyntheticsTest#pattern}
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
	// Determines whether or not the browser test variable is obfuscated.
	//
	// Can only be used with a browser variable of type `text`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#secure SyntheticsTest#secure}
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
}

