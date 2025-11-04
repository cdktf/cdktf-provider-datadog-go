// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesAmazonS3 struct {
	// The unique identifier for this component.
	//
	// Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// AWS region where the S3 bucket resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#region ObservabilityPipeline#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth *ObservabilityPipelineConfigSourcesAmazonS3Auth `field:"optional" json:"auth" yaml:"auth"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls *ObservabilityPipelineConfigSourcesAmazonS3Tls `field:"optional" json:"tls" yaml:"tls"`
}

