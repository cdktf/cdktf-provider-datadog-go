// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesGooglePubsubAuth struct {
	// Path to the GCP service account key file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/observability_pipeline#credentials_file ObservabilityPipeline#credentials_file}
	CredentialsFile *string `field:"required" json:"credentialsFile" yaml:"credentialsFile"`
}

