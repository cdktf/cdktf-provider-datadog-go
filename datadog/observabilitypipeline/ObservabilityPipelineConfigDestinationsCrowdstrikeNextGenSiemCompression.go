// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsCrowdstrikeNextGenSiemCompression struct {
	// Compression algorithm for log events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#algorithm ObservabilityPipeline#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// Compression level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#level ObservabilityPipeline#level}
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

