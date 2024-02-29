// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetManageStatusDefinition struct {
	// The query to use in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#query Dashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Whether to colorize text or background. Valid values are `background`, `text`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#color_preference Dashboard#color_preference}
	ColorPreference *string `field:"optional" json:"colorPreference" yaml:"colorPreference"`
	// The display setting to use. Valid values are `counts`, `countsAndList`, `list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#display_format Dashboard#display_format}
	DisplayFormat *string `field:"optional" json:"displayFormat" yaml:"displayFormat"`
	// A Boolean indicating whether to hide empty categories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#hide_zero_counts Dashboard#hide_zero_counts}
	HideZeroCounts interface{} `field:"optional" json:"hideZeroCounts" yaml:"hideZeroCounts"`
	// A Boolean indicating whether to show when monitors/groups last triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#show_last_triggered Dashboard#show_last_triggered}
	ShowLastTriggered interface{} `field:"optional" json:"showLastTriggered" yaml:"showLastTriggered"`
	// Whether to show the priorities column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#show_priority Dashboard#show_priority}
	ShowPriority interface{} `field:"optional" json:"showPriority" yaml:"showPriority"`
	// The method to sort the monitors.
	//
	// Valid values are `name`, `group`, `status`, `tags`, `triggered`, `group,asc`, `group,desc`, `name,asc`, `name,desc`, `status,asc`, `status,desc`, `tags,asc`, `tags,desc`, `triggered,asc`, `triggered,desc`, `priority,asc`, `priority,desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#sort Dashboard#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// The summary type to use. Valid values are `monitors`, `groups`, `combined`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#summary_type Dashboard#summary_type}
	SummaryType *string `field:"optional" json:"summaryType" yaml:"summaryType"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#title_align Dashboard#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#title_size Dashboard#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

