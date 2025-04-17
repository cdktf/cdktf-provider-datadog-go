// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSources struct {
	// datadog_agent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/observability_pipeline#datadog_agent ObservabilityPipeline#datadog_agent}
	DatadogAgent interface{} `field:"optional" json:"datadogAgent" yaml:"datadogAgent"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/observability_pipeline#kafka ObservabilityPipeline#kafka}
	Kafka interface{} `field:"optional" json:"kafka" yaml:"kafka"`
}

