// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsAmazonSecurityLake struct {
	// Name of the Amazon S3 bucket in Security Lake (3-63 characters).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#bucket ObservabilityPipeline#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Custom source name for the logs in Security Lake.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#custom_source_name ObservabilityPipeline#custom_source_name}
	CustomSourceName *string `field:"required" json:"customSourceName" yaml:"customSourceName"`
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
	// AWS region of the Security Lake bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#region ObservabilityPipeline#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth *ObservabilityPipelineConfigDestinationsAmazonSecurityLakeAuth `field:"optional" json:"auth" yaml:"auth"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls *ObservabilityPipelineConfigDestinationsAmazonSecurityLakeTls `field:"optional" json:"tls" yaml:"tls"`
}

