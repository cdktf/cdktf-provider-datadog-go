// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesOnMatchRedact struct {
	// Replacement string for redacted values (e.g., `***`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#replace ObservabilityPipeline#replace}
	Replace *string `field:"optional" json:"replace" yaml:"replace"`
}

