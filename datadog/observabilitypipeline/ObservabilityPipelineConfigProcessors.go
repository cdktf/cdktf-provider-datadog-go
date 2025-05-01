// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessors struct {
	// add_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/observability_pipeline#add_fields ObservabilityPipeline#add_fields}
	AddFields interface{} `field:"optional" json:"addFields" yaml:"addFields"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/observability_pipeline#filter ObservabilityPipeline#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// parse_json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/observability_pipeline#parse_json ObservabilityPipeline#parse_json}
	ParseJson interface{} `field:"optional" json:"parseJson" yaml:"parseJson"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/observability_pipeline#quota ObservabilityPipeline#quota}
	Quota interface{} `field:"optional" json:"quota" yaml:"quota"`
	// remove_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/observability_pipeline#remove_fields ObservabilityPipeline#remove_fields}
	RemoveFields interface{} `field:"optional" json:"removeFields" yaml:"removeFields"`
	// rename_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/observability_pipeline#rename_fields ObservabilityPipeline#rename_fields}
	RenameFields interface{} `field:"optional" json:"renameFields" yaml:"renameFields"`
}

