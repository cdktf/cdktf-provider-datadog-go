// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsOcsfMapperMapping struct {
	// Search query for selecting which logs the mapping applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// Predefined library mapping for log transformation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/observability_pipeline#library_mapping ObservabilityPipeline#library_mapping}
	LibraryMapping *string `field:"required" json:"libraryMapping" yaml:"libraryMapping"`
}

