// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsDatadogTags struct {
	// Valid values are `include`, `exclude`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#action ObservabilityPipeline#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}.
	Include *string `field:"required" json:"include" yaml:"include"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}.
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#keys ObservabilityPipeline#keys}.
	Keys *[]*string `field:"required" json:"keys" yaml:"keys"`
	// Valid values are `filter`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#mode ObservabilityPipeline#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

