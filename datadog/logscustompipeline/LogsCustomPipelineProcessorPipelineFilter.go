// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineFilter struct {
	// Filter criteria of the category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/logs_custom_pipeline#query LogsCustomPipeline#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

