// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigA struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#destinations ObservabilityPipeline#destinations}
	Destinations *ObservabilityPipelineConfigDestinations `field:"optional" json:"destinations" yaml:"destinations"`
	// processors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#processors ObservabilityPipeline#processors}
	Processors *ObservabilityPipelineConfigProcessors `field:"optional" json:"processors" yaml:"processors"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#sources ObservabilityPipeline#sources}
	Sources *ObservabilityPipelineConfigSources `field:"optional" json:"sources" yaml:"sources"`
}

