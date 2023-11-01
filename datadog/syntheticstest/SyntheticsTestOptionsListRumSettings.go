// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestOptionsListRumSettings struct {
	// Determines whether RUM data is collected during test runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/synthetics_test#is_enabled SyntheticsTest#is_enabled}
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// RUM application ID used to collect RUM data for the browser test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/synthetics_test#application_id SyntheticsTest#application_id}
	ApplicationId *string `field:"optional" json:"applicationId" yaml:"applicationId"`
	// RUM application API key ID used to collect RUM data for the browser test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/synthetics_test#client_token_id SyntheticsTest#client_token_id}
	ClientTokenId *float64 `field:"optional" json:"clientTokenId" yaml:"clientTokenId"`
}

