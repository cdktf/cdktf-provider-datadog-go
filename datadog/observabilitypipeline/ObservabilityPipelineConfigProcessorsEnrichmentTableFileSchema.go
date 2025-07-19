// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsEnrichmentTableFileSchema struct {
	// The `items` `column`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/observability_pipeline#column ObservabilityPipeline#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The type of the column (e.g. string, boolean, integer, etc.).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/observability_pipeline#type ObservabilityPipeline#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

