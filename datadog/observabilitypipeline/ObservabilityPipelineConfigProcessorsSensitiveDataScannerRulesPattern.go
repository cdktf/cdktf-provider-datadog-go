// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesPattern struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#custom ObservabilityPipeline#custom}
	Custom *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesPatternCustom `field:"optional" json:"custom" yaml:"custom"`
	// library block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#library ObservabilityPipeline#library}
	Library *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesPatternLibrary `field:"optional" json:"library" yaml:"library"`
}

