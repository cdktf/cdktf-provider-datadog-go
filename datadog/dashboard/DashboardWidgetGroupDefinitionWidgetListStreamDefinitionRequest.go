// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetListStreamDefinitionRequest struct {
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/dashboard#columns Dashboard#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/dashboard#query Dashboard#query}
	Query *DashboardWidgetGroupDefinitionWidgetListStreamDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// Widget response format. Valid values are `event_list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/dashboard#response_format Dashboard#response_format}
	ResponseFormat *string `field:"required" json:"responseFormat" yaml:"responseFormat"`
}

