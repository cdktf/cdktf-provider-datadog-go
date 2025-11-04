// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessor struct {
	// mappers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/logs_custom_pipeline#mappers LogsCustomPipeline#mappers}
	Mappers interface{} `field:"required" json:"mappers" yaml:"mappers"`
	// schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/logs_custom_pipeline#schema LogsCustomPipeline#schema}
	Schema *LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorSchema `field:"required" json:"schema" yaml:"schema"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// The name of the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

