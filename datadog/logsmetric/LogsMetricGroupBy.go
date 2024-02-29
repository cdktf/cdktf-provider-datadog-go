// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsmetric


type LogsMetricGroupBy struct {
	// The path to the value the log-based metric will be aggregated over.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/logs_metric#path LogsMetric#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Name of the tag that gets created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/logs_metric#tag_name LogsMetric#tag_name}
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
}

