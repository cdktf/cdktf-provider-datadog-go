// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipeline struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/logs_custom_pipeline#filter LogsCustomPipeline#filter}
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/logs_custom_pipeline#description LogsCustomPipeline#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/logs_custom_pipeline#processor LogsCustomPipeline#processor}
	Processor interface{} `field:"optional" json:"processor" yaml:"processor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/logs_custom_pipeline#tags LogsCustomPipeline#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

