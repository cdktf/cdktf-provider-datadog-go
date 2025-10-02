// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetScatterplotDefinition struct {
	// List of groups used for colors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#color_by_groups Dashboard#color_by_groups}
	ColorByGroups *[]*string `field:"optional" json:"colorByGroups" yaml:"colorByGroups"`
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#custom_link Dashboard#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#live_span Dashboard#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#request Dashboard#request}
	Request *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequest `field:"optional" json:"request" yaml:"request"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#title_align Dashboard#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#title_size Dashboard#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
	// xaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#xaxis Dashboard#xaxis}
	Xaxis *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionXaxis `field:"optional" json:"xaxis" yaml:"xaxis"`
	// yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#yaxis Dashboard#yaxis}
	Yaxis *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionYaxis `field:"optional" json:"yaxis" yaml:"yaxis"`
}

