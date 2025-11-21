// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorArrayProcessorOperation struct {
	// append block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/logs_custom_pipeline#append LogsCustomPipeline#append}
	Append *LogsCustomPipelineProcessorArrayProcessorOperationAppend `field:"optional" json:"append" yaml:"append"`
	// length block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/logs_custom_pipeline#length LogsCustomPipeline#length}
	Length *LogsCustomPipelineProcessorArrayProcessorOperationLength `field:"optional" json:"length" yaml:"length"`
	// select block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/logs_custom_pipeline#select LogsCustomPipeline#select}
	Select *LogsCustomPipelineProcessorArrayProcessorOperationSelect `field:"optional" json:"select" yaml:"select"`
}

