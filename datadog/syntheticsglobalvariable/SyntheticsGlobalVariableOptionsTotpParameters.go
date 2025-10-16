// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsglobalvariable


type SyntheticsGlobalVariableOptionsTotpParameters struct {
	// Number of digits for the OTP. Value must be between 4 and 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/synthetics_global_variable#digits SyntheticsGlobalVariable#digits}
	Digits *float64 `field:"required" json:"digits" yaml:"digits"`
	// Interval for which to refresh the token (in seconds). Value must be between 0 and 999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/synthetics_global_variable#refresh_interval SyntheticsGlobalVariable#refresh_interval}
	RefreshInterval *float64 `field:"required" json:"refreshInterval" yaml:"refreshInterval"`
}

