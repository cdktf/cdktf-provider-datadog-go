// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestOptionsListRetry struct {
	// Number of retries needed to consider a location as failed before sending a notification alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.29.0/docs/resources/synthetics_test#count SyntheticsTest#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Interval between a failed test and the next retry in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.29.0/docs/resources/synthetics_test#interval SyntheticsTest#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

