// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatch struct {
	// hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#hash ObservabilityPipeline#hash}
	Hash *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchHash `field:"optional" json:"hash" yaml:"hash"`
	// partial_redact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#partial_redact ObservabilityPipeline#partial_redact}
	PartialRedact *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchPartialRedact `field:"optional" json:"partialRedact" yaml:"partialRedact"`
	// redact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#redact ObservabilityPipeline#redact}
	Redact *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchRedact `field:"optional" json:"redact" yaml:"redact"`
}

