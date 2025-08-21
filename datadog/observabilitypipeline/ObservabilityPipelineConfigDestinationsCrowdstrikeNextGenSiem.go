// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsCrowdstrikeNextGenSiem struct {
	// Encoding format for log events. Valid values are `json`, `raw_message`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"required" json:"encoding" yaml:"encoding"`
	// Unique identifier for the destination component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// compression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#compression ObservabilityPipeline#compression}
	Compression *ObservabilityPipelineConfigDestinationsCrowdstrikeNextGenSiemCompression `field:"optional" json:"compression" yaml:"compression"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls *ObservabilityPipelineConfigDestinationsCrowdstrikeNextGenSiemTls `field:"optional" json:"tls" yaml:"tls"`
}

