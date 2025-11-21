// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorSchemaProcessorMappers struct {
	// schema_category_mapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/logs_custom_pipeline#schema_category_mapper LogsCustomPipeline#schema_category_mapper}
	SchemaCategoryMapper interface{} `field:"optional" json:"schemaCategoryMapper" yaml:"schemaCategoryMapper"`
	// schema_remapper block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/logs_custom_pipeline#schema_remapper LogsCustomPipeline#schema_remapper}
	SchemaRemapper interface{} `field:"optional" json:"schemaRemapper" yaml:"schemaRemapper"`
}

