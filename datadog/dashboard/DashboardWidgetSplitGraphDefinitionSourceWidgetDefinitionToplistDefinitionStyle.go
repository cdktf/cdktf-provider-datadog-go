// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinitionStyle struct {
	// display block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard#display Dashboard#display}
	Display interface{} `field:"optional" json:"display" yaml:"display"`
	// The color palette for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// The scaling mode for the widget. Valid values are `absolute`, `relative`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard#scaling Dashboard#scaling}
	Scaling *string `field:"optional" json:"scaling" yaml:"scaling"`
}

