// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesPatternCustom struct {
	// A regular expression used to detect sensitive values. Must be a valid regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/observability_pipeline#rule ObservabilityPipeline#rule}
	Rule *string `field:"optional" json:"rule" yaml:"rule"`
}

