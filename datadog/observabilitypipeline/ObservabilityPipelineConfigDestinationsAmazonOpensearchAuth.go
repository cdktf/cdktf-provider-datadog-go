// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsAmazonOpensearchAuth struct {
	// The authentication strategy to use (e.g. aws or basic).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#strategy ObservabilityPipeline#strategy}
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// ARN of the role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#assume_role ObservabilityPipeline#assume_role}
	AssumeRole *string `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// AWS region override (if applicable).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#aws_region ObservabilityPipeline#aws_region}
	AwsRegion *string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// External ID for assumed role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#external_id ObservabilityPipeline#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Session name for assumed role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#session_name ObservabilityPipeline#session_name}
	SessionName *string `field:"optional" json:"sessionName" yaml:"sessionName"`
}

