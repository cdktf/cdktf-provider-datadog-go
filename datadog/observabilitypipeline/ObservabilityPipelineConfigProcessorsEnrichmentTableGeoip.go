// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsEnrichmentTableGeoip struct {
	// Path to the IP field in the log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#key_field ObservabilityPipeline#key_field}
	KeyField *string `field:"optional" json:"keyField" yaml:"keyField"`
	// Locale used to resolve geographical names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#locale ObservabilityPipeline#locale}
	Locale *string `field:"optional" json:"locale" yaml:"locale"`
	// Path to the GeoIP database file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/observability_pipeline#path ObservabilityPipeline#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

