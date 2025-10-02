// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestinationsMicrosoftSentinel struct {
	// Azure AD client ID used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#client_id ObservabilityPipeline#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The immutable ID of the Data Collection Rule (DCR).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#dcr_immutable_id ObservabilityPipeline#dcr_immutable_id}
	DcrImmutableId *string `field:"required" json:"dcrImmutableId" yaml:"dcrImmutableId"`
	// The unique identifier for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// The name of the Log Analytics table where logs will be sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#table ObservabilityPipeline#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// Azure AD tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/observability_pipeline#tenant_id ObservabilityPipeline#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

