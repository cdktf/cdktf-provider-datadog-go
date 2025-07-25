// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestRequestClientCertificateKey struct {
	// Content of the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#content SyntheticsTest#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// File name for the certificate. Defaults to `"Provided in Terraform config"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/synthetics_test#filename SyntheticsTest#filename}
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
}

