// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorGrokParser struct {
	// grok block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/logs_custom_pipeline#grok LogsCustomPipeline#grok}
	Grok *LogsCustomPipelineProcessorPipelineProcessorGrokParserGrok `field:"required" json:"grok" yaml:"grok"`
	// Name of the log attribute to parse.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/logs_custom_pipeline#source LogsCustomPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// List of sample logs for this parser.
	//
	// It can save up to 5 samples. Each sample takes up to 5000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/logs_custom_pipeline#samples LogsCustomPipeline#samples}
	Samples *[]*string `field:"optional" json:"samples" yaml:"samples"`
}

