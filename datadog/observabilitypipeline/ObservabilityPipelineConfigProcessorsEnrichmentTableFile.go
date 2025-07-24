// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsEnrichmentTableFile struct {
	// encoding block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *ObservabilityPipelineConfigProcessorsEnrichmentTableFileEncoding `field:"optional" json:"encoding" yaml:"encoding"`
	// key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/observability_pipeline#key ObservabilityPipeline#key}
	Key interface{} `field:"optional" json:"key" yaml:"key"`
	// Path to the CSV file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/observability_pipeline#path ObservabilityPipeline#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/observability_pipeline#schema ObservabilityPipeline#schema}
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

