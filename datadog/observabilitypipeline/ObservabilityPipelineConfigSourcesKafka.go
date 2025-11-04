// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigSourcesKafka struct {
	// The Kafka consumer group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#group_id ObservabilityPipeline#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The unique ID of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of Kafka topic names to subscribe to. The source ingests messages from each topic specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#topics ObservabilityPipeline#topics}
	Topics *[]*string `field:"required" json:"topics" yaml:"topics"`
	// librdkafka_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#librdkafka_option ObservabilityPipeline#librdkafka_option}
	LibrdkafkaOption interface{} `field:"optional" json:"librdkafkaOption" yaml:"librdkafkaOption"`
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#sasl ObservabilityPipeline#sasl}
	Sasl *ObservabilityPipelineConfigSourcesKafkaSasl `field:"optional" json:"sasl" yaml:"sasl"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/observability_pipeline#tls ObservabilityPipeline#tls}
	Tls *ObservabilityPipelineConfigSourcesKafkaTls `field:"optional" json:"tls" yaml:"tls"`
}

