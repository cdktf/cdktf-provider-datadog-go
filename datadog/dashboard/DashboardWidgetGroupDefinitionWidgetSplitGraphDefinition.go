// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinition struct {
	// Size of the individual graphs in the split.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#size Dashboard#size}
	Size *string `field:"required" json:"size" yaml:"size"`
	// source_widget_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#source_widget_definition Dashboard#source_widget_definition}
	SourceWidgetDefinition *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinition `field:"required" json:"sourceWidgetDefinition" yaml:"sourceWidgetDefinition"`
	// split_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#split_config Dashboard#split_config}
	SplitConfig *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSplitConfig `field:"required" json:"splitConfig" yaml:"splitConfig"`
	// Normalize y axes across graphs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#has_uniform_y_axes Dashboard#has_uniform_y_axes}
	HasUniformYAxes interface{} `field:"optional" json:"hasUniformYAxes" yaml:"hasUniformYAxes"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#live_span Dashboard#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

