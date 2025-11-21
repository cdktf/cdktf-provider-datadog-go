// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesKeywordOptions struct {
	// A list of keywords to match near the sensitive pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/observability_pipeline#keywords ObservabilityPipeline#keywords}
	Keywords *[]*string `field:"optional" json:"keywords" yaml:"keywords"`
	// Maximum number of tokens between a keyword and a sensitive value match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/observability_pipeline#proximity ObservabilityPipeline#proximity}
	Proximity *float64 `field:"optional" json:"proximity" yaml:"proximity"`
}

