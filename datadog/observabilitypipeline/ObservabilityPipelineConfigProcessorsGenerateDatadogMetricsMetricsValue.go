// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsGenerateDatadogMetricsMetricsValue struct {
	// Metric value strategy: `increment_by_one` or `increment_by_field`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/observability_pipeline#strategy ObservabilityPipeline#strategy}
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// Name of the log field containing the numeric value to increment the metric by (used only for `increment_by_field`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/observability_pipeline#field ObservabilityPipeline#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
}

