// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsCustomProcessorRemaps struct {
	// Whether to drop events that cause errors during transformation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#drop_on_error ObservabilityPipeline#drop_on_error}
	DropOnError interface{} `field:"required" json:"dropOnError" yaml:"dropOnError"`
	// Whether this remap rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#enabled ObservabilityPipeline#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A Datadog search query used to filter events for this specific remap rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// A descriptive name for this remap rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The VRL script source code that defines the transformation logic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#source ObservabilityPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

