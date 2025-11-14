// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset


type TagPipelineRulesetRulesReferenceTableFieldPairs struct {
	// The input column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/tag_pipeline_ruleset#input_column TagPipelineRuleset#input_column}
	InputColumn *string `field:"optional" json:"inputColumn" yaml:"inputColumn"`
	// The output key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/tag_pipeline_ruleset#output_key TagPipelineRuleset#output_key}
	OutputKey *string `field:"optional" json:"outputKey" yaml:"outputKey"`
}

