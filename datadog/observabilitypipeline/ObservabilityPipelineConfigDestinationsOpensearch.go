// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsOpensearch struct {
	// The unique identifier for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of component IDs whose output is used as input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// The index or datastream to write logs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#bulk_index ObservabilityPipeline#bulk_index}
	BulkIndex *string `field:"optional" json:"bulkIndex" yaml:"bulkIndex"`
}

