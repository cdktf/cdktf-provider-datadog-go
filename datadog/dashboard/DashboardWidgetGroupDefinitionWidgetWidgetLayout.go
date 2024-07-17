// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetWidgetLayout struct {
	// The height of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#height Dashboard#height}
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// The width of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#width Dashboard#width}
	Width *float64 `field:"required" json:"width" yaml:"width"`
	// The position of the widget on the x (horizontal) axis. Must be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#x Dashboard#x}
	X *float64 `field:"required" json:"x" yaml:"x"`
	// The position of the widget on the y (vertical) axis. Must be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#y Dashboard#y}
	Y *float64 `field:"required" json:"y" yaml:"y"`
	// Whether the widget should be the first one on the second column in high density or not.
	//
	// Only one widget in the dashboard should have this property set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/dashboard#is_column_break Dashboard#is_column_break}
	IsColumnBreak interface{} `field:"optional" json:"isColumnBreak" yaml:"isColumnBreak"`
}

