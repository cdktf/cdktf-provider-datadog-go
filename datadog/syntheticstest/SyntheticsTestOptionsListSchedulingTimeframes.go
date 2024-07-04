// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestOptionsListSchedulingTimeframes struct {
	// Number representing the day of the week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#day SyntheticsTest#day}
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// The hour of the day on which scheduling starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#from SyntheticsTest#from}
	From *string `field:"required" json:"from" yaml:"from"`
	// The hour of the day on which scheduling ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/synthetics_test#to SyntheticsTest#to}
	To *string `field:"required" json:"to" yaml:"to"`
}

