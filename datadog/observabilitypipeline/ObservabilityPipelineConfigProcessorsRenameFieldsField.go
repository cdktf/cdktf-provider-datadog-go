// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsRenameFieldsField struct {
	// Destination field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/observability_pipeline#destination ObservabilityPipeline#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Whether to keep the original field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/observability_pipeline#preserve_source ObservabilityPipeline#preserve_source}
	PreserveSource interface{} `field:"required" json:"preserveSource" yaml:"preserveSource"`
	// Source field to rename.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/observability_pipeline#source ObservabilityPipeline#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

