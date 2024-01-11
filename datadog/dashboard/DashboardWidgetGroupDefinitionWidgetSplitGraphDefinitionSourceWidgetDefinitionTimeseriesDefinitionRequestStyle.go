// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyle struct {
	// The type of lines displayed. Valid values are `dashed`, `dotted`, `solid`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/dashboard#line_type Dashboard#line_type}
	LineType *string `field:"optional" json:"lineType" yaml:"lineType"`
	// The width of line displayed. Valid values are `normal`, `thick`, `thin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/dashboard#line_width Dashboard#line_width}
	LineWidth *string `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
}

