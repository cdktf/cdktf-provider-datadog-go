// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperCategories struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/logs_custom_pipeline#filter LogsCustomPipeline#filter}
	Filter *LogsCustomPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperCategoriesFilter `field:"required" json:"filter" yaml:"filter"`
	// ID to inject into the category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/logs_custom_pipeline#id LogsCustomPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"required" json:"id" yaml:"id"`
	// Value to assign to target schema field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

