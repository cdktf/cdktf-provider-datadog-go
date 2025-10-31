// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awscurconfig


type AwsCurConfigAccountFilters struct {
	// List of AWS account IDs to exclude from cost analysis.
	//
	// Only used when `include_new_accounts` is `true`. Cannot be used together with `included_accounts`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/aws_cur_config#excluded_accounts AwsCurConfig#excluded_accounts}
	ExcludedAccounts *[]*string `field:"optional" json:"excludedAccounts" yaml:"excludedAccounts"`
	// List of AWS account IDs to include in cost analysis.
	//
	// Only used when `include_new_accounts` is `false`. Cannot be used together with `excluded_accounts`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/aws_cur_config#included_accounts AwsCurConfig#included_accounts}
	IncludedAccounts *[]*string `field:"optional" json:"includedAccounts" yaml:"includedAccounts"`
	// Whether to automatically include new member accounts in your cost analysis.
	//
	// When `true`, use `excluded_accounts` to specify accounts to exclude. When `false`, use `included_accounts` to specify only the accounts to include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/aws_cur_config#include_new_accounts AwsCurConfig#include_new_accounts}
	IncludeNewAccounts interface{} `field:"optional" json:"includeNewAccounts" yaml:"includeNewAccounts"`
}

