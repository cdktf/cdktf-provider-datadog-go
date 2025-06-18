// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesScope struct {
	// Scan all fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#all ObservabilityPipeline#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// exclude block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#exclude ObservabilityPipeline#exclude}
	Exclude *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesScopeExclude `field:"optional" json:"exclude" yaml:"exclude"`
	// include block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesScopeInclude `field:"optional" json:"include" yaml:"include"`
}

