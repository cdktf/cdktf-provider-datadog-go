// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesPatternLibrary struct {
	// Identifier for a predefined pattern from the sensitive data scanner pattern library.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to augment the pattern with recommended keywords (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#use_recommended_keywords ObservabilityPipeline#use_recommended_keywords}
	UseRecommendedKeywords interface{} `field:"optional" json:"useRecommendedKeywords" yaml:"useRecommendedKeywords"`
}

