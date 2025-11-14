// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRules struct {
	// keyword_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#keyword_options ObservabilityPipeline#keyword_options}
	KeywordOptions *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesKeywordOptions `field:"optional" json:"keywordOptions" yaml:"keywordOptions"`
	// A name identifying the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// on_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#on_match ObservabilityPipeline#on_match}
	OnMatch *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatch `field:"optional" json:"onMatch" yaml:"onMatch"`
	// pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#pattern ObservabilityPipeline#pattern}
	Pattern *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesPattern `field:"optional" json:"pattern" yaml:"pattern"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#scope ObservabilityPipeline#scope}
	Scope *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesScope `field:"optional" json:"scope" yaml:"scope"`
	// Tags assigned to this rule for filtering and classification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/observability_pipeline#tags ObservabilityPipeline#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

