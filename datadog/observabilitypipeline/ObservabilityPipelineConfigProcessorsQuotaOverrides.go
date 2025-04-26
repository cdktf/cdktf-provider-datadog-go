// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsQuotaOverrides struct {
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/observability_pipeline#limit ObservabilityPipeline#limit}
	Limit *ObservabilityPipelineConfigProcessorsQuotaOverridesLimit `field:"required" json:"limit" yaml:"limit"`
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/observability_pipeline#field ObservabilityPipeline#field}
	Field interface{} `field:"optional" json:"field" yaml:"field"`
}

