// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestApiStepExtractedValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/synthetics_test#name SyntheticsTest#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/synthetics_test#parser SyntheticsTest#parser}
	Parser *SyntheticsTestApiStepExtractedValueParser `field:"required" json:"parser" yaml:"parser"`
	// Property of the Synthetics Test Response to use for the variable. Valid values are `grpc_message`, `grpc_metadata`, `http_body`, `http_header`, `http_status_code`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// When type is `http_header` or `grpc_metadata`, name of the header or metadatum to extract.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/synthetics_test#field SyntheticsTest#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Determines whether or not the extracted value will be obfuscated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/synthetics_test#secure SyntheticsTest#secure}
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
}

