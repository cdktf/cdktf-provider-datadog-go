// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorSchemaProcessorSchema struct {
	// Class name of the schema to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#class_name LogsCustomPipeline#class_name}
	ClassName *string `field:"required" json:"className" yaml:"className"`
	// Class UID of the schema to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#class_uid LogsCustomPipeline#class_uid}
	ClassUid *float64 `field:"required" json:"classUid" yaml:"classUid"`
	// Type of schema to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#schema_type LogsCustomPipeline#schema_type}
	SchemaType *string `field:"required" json:"schemaType" yaml:"schemaType"`
	// Version of the schema to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#version LogsCustomPipeline#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Optional list of extensions to modify the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#extensions LogsCustomPipeline#extensions}
	Extensions *[]*string `field:"optional" json:"extensions" yaml:"extensions"`
	// Optional list of profiles to modify the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#profiles LogsCustomPipeline#profiles}
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
}

