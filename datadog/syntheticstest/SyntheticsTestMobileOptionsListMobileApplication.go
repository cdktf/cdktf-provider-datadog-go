// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileOptionsListMobileApplication struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/synthetics_test#application_id SyntheticsTest#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/synthetics_test#reference_id SyntheticsTest#reference_id}.
	ReferenceId *string `field:"required" json:"referenceId" yaml:"referenceId"`
	// Valid values are `latest`, `version`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/synthetics_test#reference_type SyntheticsTest#reference_type}
	ReferenceType *string `field:"required" json:"referenceType" yaml:"referenceType"`
}

