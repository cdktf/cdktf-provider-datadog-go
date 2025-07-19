// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsEnrichmentTableFileEncoding struct {
	// The `encoding` `delimiter`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/observability_pipeline#delimiter ObservabilityPipeline#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// The `encoding` `includes_headers`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/observability_pipeline#includes_headers ObservabilityPipeline#includes_headers}
	IncludesHeaders interface{} `field:"optional" json:"includesHeaders" yaml:"includesHeaders"`
	// File encoding format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/observability_pipeline#type ObservabilityPipeline#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

