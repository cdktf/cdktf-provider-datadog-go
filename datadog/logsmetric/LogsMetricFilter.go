// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsmetric


type LogsMetricFilter struct {
	// The search query - following the log search syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/logs_metric#query LogsMetric#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

