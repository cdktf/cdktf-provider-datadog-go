// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannerrule


type SensitiveDataScannerRuleIncludedKeywordConfiguration struct {
	// Number of characters before the match to find a keyword validating the match.
	//
	// It must be between 1 and 50 (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/sensitive_data_scanner_rule#character_count SensitiveDataScannerRule#character_count}
	CharacterCount *float64 `field:"required" json:"characterCount" yaml:"characterCount"`
	// Keyword list that is checked during scanning in order to validate a match.
	//
	// The number of keywords in the list must be lower than or equal to 30.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/sensitive_data_scanner_rule#keywords SensitiveDataScannerRule#keywords}
	Keywords *[]*string `field:"required" json:"keywords" yaml:"keywords"`
}

