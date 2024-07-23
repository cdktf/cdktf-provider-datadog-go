// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsglobalvariable


type SyntheticsGlobalVariableOptionsTotpParameters struct {
	// Number of digits for the OTP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/synthetics_global_variable#digits SyntheticsGlobalVariable#digits}
	Digits *float64 `field:"required" json:"digits" yaml:"digits"`
	// Interval for which to refresh the token (in seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/synthetics_global_variable#refresh_interval SyntheticsGlobalVariable#refresh_interval}
	RefreshInterval *float64 `field:"required" json:"refreshInterval" yaml:"refreshInterval"`
}

