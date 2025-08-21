// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsParseGrokRules struct {
	// The name of the field in the log event to apply the Grok rules to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#source ObservabilityPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// match_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#match_rule ObservabilityPipeline#match_rule}
	MatchRule interface{} `field:"optional" json:"matchRule" yaml:"matchRule"`
	// support_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#support_rule ObservabilityPipeline#support_rule}
	SupportRule interface{} `field:"optional" json:"supportRule" yaml:"supportRule"`
}

