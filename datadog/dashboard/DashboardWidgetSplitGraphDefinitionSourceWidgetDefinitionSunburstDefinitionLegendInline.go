// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionLegendInline struct {
	// The type of legend (inline or automatic). Valid values are `inline`, `automatic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether to hide the percentages of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#hide_percent Dashboard#hide_percent}
	HidePercent interface{} `field:"optional" json:"hidePercent" yaml:"hidePercent"`
	// Whether to hide the values of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#hide_value Dashboard#hide_value}
	HideValue interface{} `field:"optional" json:"hideValue" yaml:"hideValue"`
}

