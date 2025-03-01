// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestMobileOptionsListBindings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/synthetics_test#principals SyntheticsTest#principals}.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
	// Valid values are `editor`, `viewer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/synthetics_test#relation SyntheticsTest#relation}
	Relation *string `field:"optional" json:"relation" yaml:"relation"`
}

