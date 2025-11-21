// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesSocketFramingCharacterDelimited struct {
	// A single ASCII character used as a delimiter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/observability_pipeline#delimiter ObservabilityPipeline#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

