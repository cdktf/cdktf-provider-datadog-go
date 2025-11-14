// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsindex


type LogsIndexExclusionFilterFilter struct {
	// Only logs matching the filter criteria and the query of the parent index will be considered for this exclusion filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_index#query LogsIndex#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// The fraction of logs excluded by the exclusion filter, when active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_index#sample_rate LogsIndex#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

