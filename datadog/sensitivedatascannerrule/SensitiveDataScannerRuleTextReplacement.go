// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannerrule


type SensitiveDataScannerRuleTextReplacement struct {
	// Type of the replacement text.
	//
	// None means no replacement. hash means the data will be stubbed. replacement_string means that one can chose a text to replace the data. partial_replacement_from_beginning allows a user to partially replace the data from the beginning, and partial_replacement_from_end on the other hand, allows to replace data from the end. Valid values are `none`, `hash`, `replacement_string`, `partial_replacement_from_beginning`, `partial_replacement_from_end`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/sensitive_data_scanner_rule#type SensitiveDataScannerRule#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Required if type == 'partial_replacement_from_beginning' or 'partial_replacement_from_end'. It must be > 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/sensitive_data_scanner_rule#number_of_chars SensitiveDataScannerRule#number_of_chars}
	NumberOfChars *float64 `field:"optional" json:"numberOfChars" yaml:"numberOfChars"`
	// Required if type == 'replacement_string'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/sensitive_data_scanner_rule#replacement_string SensitiveDataScannerRule#replacement_string}
	ReplacementString *string `field:"optional" json:"replacementString" yaml:"replacementString"`
}

