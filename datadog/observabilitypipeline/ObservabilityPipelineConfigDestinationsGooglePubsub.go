// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsGooglePubsub struct {
	// The unique identifier for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// The GCP project ID that owns the Pub/Sub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#project ObservabilityPipeline#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The Pub/Sub topic name to publish logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#topic ObservabilityPipeline#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth *ObservabilityPipelineConfigDestinationsGooglePubsubAuth `field:"optional" json:"auth" yaml:"auth"`
	// Encoding format for log events. Valid values: `json`, `raw_message`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls *ObservabilityPipelineConfigDestinationsGooglePubsubTls `field:"optional" json:"tls" yaml:"tls"`
}

