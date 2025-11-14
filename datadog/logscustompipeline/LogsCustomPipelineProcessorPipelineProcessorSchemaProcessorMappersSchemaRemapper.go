// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorPipelineProcessorSchemaProcessorMappersSchemaRemapper struct {
	// Name of the logs schema remapper.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Array of source attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#sources LogsCustomPipeline#sources}
	Sources *[]*string `field:"required" json:"sources" yaml:"sources"`
	// Target field to map log source field to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Override or not the target element if already set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#override_on_conflict LogsCustomPipeline#override_on_conflict}
	OverrideOnConflict interface{} `field:"optional" json:"overrideOnConflict" yaml:"overrideOnConflict"`
	// Remove or preserve the remapped source element.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#preserve_source LogsCustomPipeline#preserve_source}
	PreserveSource interface{} `field:"optional" json:"preserveSource" yaml:"preserveSource"`
	// If the `target_type` of the remapper is `attribute`, try to cast the value to a new specific type.
	//
	// If the cast is not possible, the original type is kept. `string`, `integer`, or `double` are the possible types. If the `target_type` is `tag`, this parameter may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/logs_custom_pipeline#target_format LogsCustomPipeline#target_format}
	TargetFormat *string `field:"optional" json:"targetFormat" yaml:"targetFormat"`
}

