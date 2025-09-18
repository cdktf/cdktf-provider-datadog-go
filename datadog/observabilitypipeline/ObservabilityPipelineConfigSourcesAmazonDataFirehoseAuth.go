// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesAmazonDataFirehoseAuth struct {
	// The Amazon Resource Name (ARN) of the role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/observability_pipeline#assume_role ObservabilityPipeline#assume_role}
	AssumeRole *string `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// A unique identifier for cross-account role assumption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/observability_pipeline#external_id ObservabilityPipeline#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// A session identifier used for logging and tracing the assumed role session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/observability_pipeline#session_name ObservabilityPipeline#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
}

