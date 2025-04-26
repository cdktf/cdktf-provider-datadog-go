// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinations struct {
	// datadog_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/observability_pipeline#datadog_logs ObservabilityPipeline#datadog_logs}
	DatadogLogs interface{} `field:"optional" json:"datadogLogs" yaml:"datadogLogs"`
}

