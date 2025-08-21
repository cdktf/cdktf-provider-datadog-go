// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsGoogleChronicle struct {
	// The unique identifier for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#auth ObservabilityPipeline#auth}
	Auth *ObservabilityPipelineConfigDestinationsGoogleChronicleAuth `field:"optional" json:"auth" yaml:"auth"`
	// The Google Chronicle customer ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#customer_id ObservabilityPipeline#customer_id}
	CustomerId *string `field:"optional" json:"customerId" yaml:"customerId"`
	// The encoding format for the logs sent to Chronicle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#encoding ObservabilityPipeline#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The log type metadata associated with the Chronicle destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/observability_pipeline#log_type ObservabilityPipeline#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

