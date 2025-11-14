// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapper struct {
	// categories block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#categories LogsCustomPipeline#categories}
	Categories interface{} `field:"required" json:"categories" yaml:"categories"`
	// Name of the logs schema category mapper.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#targets LogsCustomPipeline#targets}
	Targets *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperTargets `field:"required" json:"targets" yaml:"targets"`
	// fallback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#fallback LogsCustomPipeline#fallback}
	Fallback *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback `field:"optional" json:"fallback" yaml:"fallback"`
}

