// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apmretentionfilter


type ApmRetentionFilterFilter struct {
	// The search query - follow the span search syntax, use `AND` between tags and `\` to escape special characters, use nanosecond for duration.
	//
	// Defaults to `"*"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/apm_retention_filter#query ApmRetentionFilter#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

