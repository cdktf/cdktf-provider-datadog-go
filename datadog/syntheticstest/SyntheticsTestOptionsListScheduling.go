// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestOptionsListScheduling struct {
	// timeframes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/synthetics_test#timeframes SyntheticsTest#timeframes}
	Timeframes interface{} `field:"required" json:"timeframes" yaml:"timeframes"`
	// Timezone in which the timeframe is based.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/synthetics_test#timezone SyntheticsTest#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

