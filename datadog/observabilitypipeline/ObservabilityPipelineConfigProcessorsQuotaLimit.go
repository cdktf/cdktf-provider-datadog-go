// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsQuotaLimit struct {
	// Whether to enforce by 'bytes' or 'events'. Valid values are `bytes`, `events`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#enforce ObservabilityPipeline#enforce}
	Enforce *string `field:"required" json:"enforce" yaml:"enforce"`
	// The daily quota limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#limit ObservabilityPipeline#limit}
	Limit *float64 `field:"required" json:"limit" yaml:"limit"`
}

