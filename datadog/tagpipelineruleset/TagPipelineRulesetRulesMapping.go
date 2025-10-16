// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagpipelineruleset


type TagPipelineRulesetRulesMapping struct {
	// The destination key for the mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#destination_key TagPipelineRuleset#destination_key}
	DestinationKey *string `field:"optional" json:"destinationKey" yaml:"destinationKey"`
	// Whether to apply the mapping only if the destination key doesn't exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#if_not_exists TagPipelineRuleset#if_not_exists}
	IfNotExists interface{} `field:"optional" json:"ifNotExists" yaml:"ifNotExists"`
	// The source keys for the mapping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/tag_pipeline_ruleset#source_keys TagPipelineRuleset#source_keys}
	SourceKeys *[]*string `field:"optional" json:"sourceKeys" yaml:"sourceKeys"`
}

