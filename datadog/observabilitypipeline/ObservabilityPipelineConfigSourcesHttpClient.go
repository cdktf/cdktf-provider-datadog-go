// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesHttpClient struct {
	// The decoding format used to interpret incoming logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#decoding ObservabilityPipeline#decoding}
	Decoding *string `field:"required" json:"decoding" yaml:"decoding"`
	// The unique identifier for this component.
	//
	// Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Optional authentication strategy for HTTP requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#auth_strategy ObservabilityPipeline#auth_strategy}
	AuthStrategy *string `field:"optional" json:"authStrategy" yaml:"authStrategy"`
	// The interval (in seconds) between HTTP scrape requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#scrape_interval_secs ObservabilityPipeline#scrape_interval_secs}
	ScrapeIntervalSecs *float64 `field:"optional" json:"scrapeIntervalSecs" yaml:"scrapeIntervalSecs"`
	// The timeout (in seconds) for each scrape request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#scrape_timeout_secs ObservabilityPipeline#scrape_timeout_secs}
	ScrapeTimeoutSecs *float64 `field:"optional" json:"scrapeTimeoutSecs" yaml:"scrapeTimeoutSecs"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls *ObservabilityPipelineConfigSourcesHttpClientTls `field:"optional" json:"tls" yaml:"tls"`
}

