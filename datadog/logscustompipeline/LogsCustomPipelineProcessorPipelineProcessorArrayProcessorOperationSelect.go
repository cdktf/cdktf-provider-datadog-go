// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorArrayProcessorOperationSelect struct {
	// Filter expression (e.g. key1:value1 OR key2:value2) used to find the matching element.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/logs_custom_pipeline#filter LogsCustomPipeline#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Attribute path of the array to search into.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/logs_custom_pipeline#source LogsCustomPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Attribute that receives the extracted value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Attribute key from the matching object that should be extracted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/logs_custom_pipeline#value_to_extract LogsCustomPipeline#value_to_extract}
	ValueToExtract *string `field:"required" json:"valueToExtract" yaml:"valueToExtract"`
}

