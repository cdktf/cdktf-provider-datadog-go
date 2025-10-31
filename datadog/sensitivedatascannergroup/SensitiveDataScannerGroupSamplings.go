// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannergroup


type SensitiveDataScannerGroupSamplings struct {
	// Product that the sampling rate applies to. Valid values are `logs`, `rum`, `events`, `apm`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/sensitive_data_scanner_group#product SensitiveDataScannerGroup#product}
	Product *string `field:"required" json:"product" yaml:"product"`
	// Percentage rate at which data for the product type is scanned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/sensitive_data_scanner_group#rate SensitiveDataScannerGroup#rate}
	Rate *float64 `field:"required" json:"rate" yaml:"rate"`
}

