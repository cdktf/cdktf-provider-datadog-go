// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsReduceMergeStrategies struct {
	// The field path in the log event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#path ObservabilityPipeline#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The merge strategy to apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#strategy ObservabilityPipeline#strategy}
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
}

