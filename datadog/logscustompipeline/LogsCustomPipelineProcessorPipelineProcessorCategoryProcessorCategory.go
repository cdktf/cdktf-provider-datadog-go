// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorCategory struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/logs_custom_pipeline#filter LogsCustomPipeline#filter}
	Filter *LogsCustomPipelineProcessorPipelineProcessorCategoryProcessorCategoryFilter `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

