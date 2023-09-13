// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spansmetric


type SpansMetricCompute struct {
	// The type of aggregation to use. This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/spans_metric#aggregation_type SpansMetric#aggregation_type}
	AggregationType *string `field:"required" json:"aggregationType" yaml:"aggregationType"`
	// Toggle to include or exclude percentile aggregations for distribution metrics. Only present when the `aggregation_type` is `distribution`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/spans_metric#include_percentiles SpansMetric#include_percentiles}
	IncludePercentiles interface{} `field:"optional" json:"includePercentiles" yaml:"includePercentiles"`
	// The path to the value the span-based metric will aggregate on (only used if the aggregation type is a "distribution").
	//
	// This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/spans_metric#path SpansMetric#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

