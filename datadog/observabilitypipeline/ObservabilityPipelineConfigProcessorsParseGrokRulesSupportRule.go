// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsParseGrokRulesSupportRule struct {
	// The name of the helper Grok rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The definition of the helper Grok rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#rule ObservabilityPipeline#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
}

