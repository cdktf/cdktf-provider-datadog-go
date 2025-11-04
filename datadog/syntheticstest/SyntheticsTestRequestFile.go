// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestRequestFile struct {
	// Name of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/synthetics_test#name SyntheticsTest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Size of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/synthetics_test#size SyntheticsTest#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Type of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/synthetics_test#type SyntheticsTest#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Content of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/synthetics_test#content SyntheticsTest#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Original name of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/synthetics_test#original_file_name SyntheticsTest#original_file_name}
	OriginalFileName *string `field:"optional" json:"originalFileName" yaml:"originalFileName"`
}

