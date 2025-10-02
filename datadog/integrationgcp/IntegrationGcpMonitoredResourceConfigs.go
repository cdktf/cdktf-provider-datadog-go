// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationgcp


type IntegrationGcpMonitoredResourceConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/integration_gcp#filters IntegrationGcp#filters}.
	Filters *[]*string `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/integration_gcp#type IntegrationGcp#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

