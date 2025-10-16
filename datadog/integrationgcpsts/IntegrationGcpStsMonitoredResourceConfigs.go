// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationgcpsts


type IntegrationGcpStsMonitoredResourceConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/integration_gcp_sts#filters IntegrationGcpSts#filters}.
	Filters *[]*string `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/integration_gcp_sts#type IntegrationGcpSts#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

