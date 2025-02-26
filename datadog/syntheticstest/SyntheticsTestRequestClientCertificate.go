// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestRequestClientCertificate struct {
	// cert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#cert SyntheticsTest#cert}
	Cert *SyntheticsTestRequestClientCertificateCert `field:"required" json:"cert" yaml:"cert"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/synthetics_test#key SyntheticsTest#key}
	Key *SyntheticsTestRequestClientCertificateKey `field:"required" json:"key" yaml:"key"`
}

