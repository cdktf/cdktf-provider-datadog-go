// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorAttributeRemapper struct {
	// List of source attributes or tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#sources LogsCustomPipeline#sources}
	Sources *[]*string `field:"required" json:"sources" yaml:"sources"`
	// Defines where the sources are from (log `attribute` or `tag`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#source_type LogsCustomPipeline#source_type}
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Final attribute or tag name to remap the sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// Defines if the target is a log `attribute` or `tag`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#target_type LogsCustomPipeline#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Override the target element if already set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#override_on_conflict LogsCustomPipeline#override_on_conflict}
	OverrideOnConflict interface{} `field:"optional" json:"overrideOnConflict" yaml:"overrideOnConflict"`
	// Remove or preserve the remapped source element.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#preserve_source LogsCustomPipeline#preserve_source}
	PreserveSource interface{} `field:"optional" json:"preserveSource" yaml:"preserveSource"`
	// If the `target_type` of the remapper is `attribute`, try to cast the value to a new specific type.
	//
	// If the cast is not possible, the original type is kept. `string`, `integer`, or `double` are the possible types. If the `target_type` is `tag`, this parameter may not be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/logs_custom_pipeline#target_format LogsCustomPipeline#target_format}
	TargetFormat *string `field:"optional" json:"targetFormat" yaml:"targetFormat"`
}

