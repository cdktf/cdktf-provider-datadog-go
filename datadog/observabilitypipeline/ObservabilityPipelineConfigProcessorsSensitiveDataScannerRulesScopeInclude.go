// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsSensitiveDataScannerRulesScopeInclude struct {
	// The fields to include in scanning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/observability_pipeline#fields ObservabilityPipeline#fields}
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
}

