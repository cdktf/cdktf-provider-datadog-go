// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchPartialRedact struct {
	// Number of characters to keep.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#characters ObservabilityPipeline#characters}
	Characters *float64 `field:"optional" json:"characters" yaml:"characters"`
	// Direction from which to keep characters: `first` or `last`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#direction ObservabilityPipeline#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
}

