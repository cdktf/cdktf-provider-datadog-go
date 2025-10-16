// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset


type TagPipelineRulesetRulesReferenceTable struct {
	// Whether the reference table lookup is case insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#case_insensitivity TagPipelineRuleset#case_insensitivity}
	CaseInsensitivity interface{} `field:"optional" json:"caseInsensitivity" yaml:"caseInsensitivity"`
	// field_pairs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#field_pairs TagPipelineRuleset#field_pairs}
	FieldPairs interface{} `field:"optional" json:"fieldPairs" yaml:"fieldPairs"`
	// Whether to apply the reference table only if the key doesn't exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#if_not_exists TagPipelineRuleset#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// The source keys for the reference table lookup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#source_keys TagPipelineRuleset#source_keys}
	SourceKeys *[]*string `field:"optional" json:"sourceKeys" yaml:"sourceKeys"`
	// The name of the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#table_name TagPipelineRuleset#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

