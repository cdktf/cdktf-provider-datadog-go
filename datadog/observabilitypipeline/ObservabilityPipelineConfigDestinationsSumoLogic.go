// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsSumoLogic struct {
	// The unique identifier for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// The output encoding format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// header_custom_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#header_custom_fields ObservabilityPipeline#header_custom_fields}
	HeaderCustomFields interface{} `field:"optional" json:"headerCustomFields" yaml:"headerCustomFields"`
	// Optional override for the host name header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#header_host_name ObservabilityPipeline#header_host_name}
	HeaderHostName *string `field:"optional" json:"headerHostName" yaml:"headerHostName"`
	// Optional override for the source category header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#header_source_category ObservabilityPipeline#header_source_category}
	HeaderSourceCategory *string `field:"optional" json:"headerSourceCategory" yaml:"headerSourceCategory"`
	// Optional override for the source name header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#header_source_name ObservabilityPipeline#header_source_name}
	HeaderSourceName *string `field:"optional" json:"headerSourceName" yaml:"headerSourceName"`
}

