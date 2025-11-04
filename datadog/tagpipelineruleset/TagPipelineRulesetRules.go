// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset


type TagPipelineRulesetRules struct {
	// Whether the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/tag_pipeline_ruleset#enabled TagPipelineRuleset#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The name of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/tag_pipeline_ruleset#name TagPipelineRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/tag_pipeline_ruleset#mapping TagPipelineRuleset#mapping}
	Mapping *TagPipelineRulesetRulesMapping `field:"optional" json:"mapping" yaml:"mapping"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/tag_pipeline_ruleset#query TagPipelineRuleset#query}
	Query *TagPipelineRulesetRulesQuery `field:"optional" json:"query" yaml:"query"`
	// reference_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/tag_pipeline_ruleset#reference_table TagPipelineRuleset#reference_table}
	ReferenceTable *TagPipelineRulesetRulesReferenceTable `field:"optional" json:"referenceTable" yaml:"referenceTable"`
}

