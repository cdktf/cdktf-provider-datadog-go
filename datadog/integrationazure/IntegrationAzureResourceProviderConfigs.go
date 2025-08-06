// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationazure


type IntegrationAzureResourceProviderConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/integration_azure#metrics_enabled IntegrationAzure#metrics_enabled}.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/integration_azure#namespace IntegrationAzure#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

