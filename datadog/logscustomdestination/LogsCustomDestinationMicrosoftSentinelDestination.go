// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logscustomdestination


type LogsCustomDestinationMicrosoftSentinelDestination struct {
	// Client ID from the Datadog Azure Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_custom_destination#client_id LogsCustomDestination#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Azure Data Collection Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_custom_destination#data_collection_endpoint LogsCustomDestination#data_collection_endpoint}
	DataCollectionEndpoint *string `field:"required" json:"dataCollectionEndpoint" yaml:"dataCollectionEndpoint"`
	// Azure Data Collection Rule ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_custom_destination#data_collection_rule_id LogsCustomDestination#data_collection_rule_id}
	DataCollectionRuleId *string `field:"required" json:"dataCollectionRuleId" yaml:"dataCollectionRuleId"`
	// Azure stream name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_custom_destination#stream_name LogsCustomDestination#stream_name}
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// Tenant ID from the Datadog Azure Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_custom_destination#tenant_id LogsCustomDestination#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

