// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSloListDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#query Dashboard#query}
	Query *DashboardWidgetSloListDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// The request type for the SLO List request. Valid values are `slo_list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#request_type Dashboard#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
}

