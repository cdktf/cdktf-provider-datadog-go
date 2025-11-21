// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaCategoryMapperFallback struct {
	// Fallback sources used to populate value of field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/logs_custom_pipeline#sources LogsCustomPipeline#sources}
	Sources *map[string]*string `field:"optional" json:"sources" yaml:"sources"`
	// Values that define when the fallback is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/logs_custom_pipeline#values LogsCustomPipeline#values}
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}

