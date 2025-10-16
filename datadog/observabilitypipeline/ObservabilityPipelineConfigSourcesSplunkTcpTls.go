// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesSplunkTcpTls struct {
	// Path to the Certificate Authority (CA) file used to validate the server's TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/observability_pipeline#ca_file ObservabilityPipeline#ca_file}
	CaFile *string `field:"optional" json:"caFile" yaml:"caFile"`
	// Path to the TLS client certificate file used to authenticate the pipeline component with upstream or downstream services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/observability_pipeline#crt_file ObservabilityPipeline#crt_file}
	CrtFile *string `field:"optional" json:"crtFile" yaml:"crtFile"`
	// Path to the private key file associated with the TLS client certificate. Used for mutual TLS authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/observability_pipeline#key_file ObservabilityPipeline#key_file}
	KeyFile *string `field:"optional" json:"keyFile" yaml:"keyFile"`
}

