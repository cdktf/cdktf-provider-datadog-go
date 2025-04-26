// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesKafkaSasl struct {
	// SASL mechanism to use (e.g., PLAIN, SCRAM-SHA-256, SCRAM-SHA-512). Valid values are `PLAIN`, `SCRAM-SHA-256`, `SCRAM-SHA-512`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/observability_pipeline#mechanism ObservabilityPipeline#mechanism}
	Mechanism *string `field:"required" json:"mechanism" yaml:"mechanism"`
}

