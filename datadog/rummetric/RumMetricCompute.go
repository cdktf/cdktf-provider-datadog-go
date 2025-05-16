// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rummetric


type RumMetricCompute struct {
	// The type of aggregation to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/rum_metric#aggregation_type RumMetric#aggregation_type}
	AggregationType *string `field:"required" json:"aggregationType" yaml:"aggregationType"`
	// Toggle to include or exclude percentile aggregations for distribution metrics. Only present when `aggregation_type` is `distribution`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/rum_metric#include_percentiles RumMetric#include_percentiles}
	IncludePercentiles interface{} `field:"optional" json:"includePercentiles" yaml:"includePercentiles"`
	// The path to the value the RUM-based metric will aggregate on. Only present when `aggregation_type` is `distribution`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/rum_metric#path RumMetric#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

