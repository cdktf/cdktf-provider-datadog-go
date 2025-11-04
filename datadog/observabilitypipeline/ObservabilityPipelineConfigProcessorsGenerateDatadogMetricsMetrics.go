// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsGenerateDatadogMetricsMetrics struct {
	// Datadog filter query to match logs for metric generation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// Type of metric to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#metric_type ObservabilityPipeline#metric_type}
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// Name of the custom metric to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional fields used to group the metric series.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#group_by ObservabilityPipeline#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#value ObservabilityPipeline#value}
	Value *ObservabilityPipelineConfigProcessorsGenerateDatadogMetricsMetricsValue `field:"optional" json:"value" yaml:"value"`
}

