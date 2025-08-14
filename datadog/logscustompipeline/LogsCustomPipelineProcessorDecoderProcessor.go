// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustompipeline


type LogsCustomPipelineProcessorDecoderProcessor struct {
	// Encoding type: base64 or base16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_pipeline#binary_to_text_encoding LogsCustomPipeline#binary_to_text_encoding}
	BinaryToTextEncoding *string `field:"required" json:"binaryToTextEncoding" yaml:"binaryToTextEncoding"`
	// Input representation: utf-8 or integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_pipeline#input_representation LogsCustomPipeline#input_representation}
	InputRepresentation *string `field:"required" json:"inputRepresentation" yaml:"inputRepresentation"`
	// Encoded message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_pipeline#source LogsCustomPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Decoded message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_pipeline#target LogsCustomPipeline#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// If the processor is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_pipeline#is_enabled LogsCustomPipeline#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Name of the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/logs_custom_pipeline#name LogsCustomPipeline#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

