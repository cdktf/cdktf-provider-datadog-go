// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsDedupe struct {
	// A list of log field paths to check for duplicates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#fields ObservabilityPipeline#fields}
	Fields *[]*string `field:"required" json:"fields" yaml:"fields"`
	// The unique identifier for this processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// A list of component IDs whose output is used as the input for this processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// The deduplication mode to apply to the fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#mode ObservabilityPipeline#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

