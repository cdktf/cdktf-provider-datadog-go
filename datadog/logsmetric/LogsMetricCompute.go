// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsmetric


type LogsMetricCompute struct {
	// The type of aggregation to use. This field can't be updated after creation. Valid values are `count`, `distribution`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/logs_metric#aggregation_type LogsMetric#aggregation_type}
	AggregationType *string `field:"required" json:"aggregationType" yaml:"aggregationType"`
	// Toggle to include/exclude percentiles for a distribution metric.
	//
	// Defaults to false. Can only be applied to metrics that have an `aggregation_type` of distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/logs_metric#include_percentiles LogsMetric#include_percentiles}
	IncludePercentiles interface{} `field:"optional" json:"includePercentiles" yaml:"includePercentiles"`
	// The path to the value the log-based metric will aggregate on (only used if the aggregation type is a "distribution").
	//
	// This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/logs_metric#path LogsMetric#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

