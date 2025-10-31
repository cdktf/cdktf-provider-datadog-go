// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset


type TagPipelineRulesetRulesQueryAddition struct {
	// The key to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/tag_pipeline_ruleset#key TagPipelineRuleset#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/tag_pipeline_ruleset#value TagPipelineRuleset#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

