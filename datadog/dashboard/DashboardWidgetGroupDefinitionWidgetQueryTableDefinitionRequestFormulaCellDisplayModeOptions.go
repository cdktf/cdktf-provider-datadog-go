// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions struct {
	// The type of trend line to display. Valid values are `area`, `line`, `bars`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#trend_type Dashboard#trend_type}
	TrendType *string `field:"optional" json:"trendType" yaml:"trendType"`
	// The scale of the y-axis. Valid values are `shared`, `independent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#y_scale Dashboard#y_scale}
	YScale *string `field:"optional" json:"yScale" yaml:"yScale"`
}

