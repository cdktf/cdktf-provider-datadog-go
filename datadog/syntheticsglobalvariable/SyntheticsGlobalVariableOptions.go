// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsglobalvariable


type SyntheticsGlobalVariableOptions struct {
	// totp_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/synthetics_global_variable#totp_parameters SyntheticsGlobalVariable#totp_parameters}
	TotpParameters interface{} `field:"optional" json:"totpParameters" yaml:"totpParameters"`
}

