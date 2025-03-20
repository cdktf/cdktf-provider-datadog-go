// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetManageStatusDefinition struct {
	// The query to use in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#query Powerpack#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Whether to colorize text or background. Valid values are `background`, `text`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#color_preference Powerpack#color_preference}
	ColorPreference *string `field:"optional" json:"colorPreference" yaml:"colorPreference"`
	// The display setting to use. Valid values are `counts`, `countsAndList`, `list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#display_format Powerpack#display_format}
	DisplayFormat *string `field:"optional" json:"displayFormat" yaml:"displayFormat"`
	// A Boolean indicating whether to hide empty categories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#hide_zero_counts Powerpack#hide_zero_counts}
	HideZeroCounts interface{} `field:"optional" json:"hideZeroCounts" yaml:"hideZeroCounts"`
	// A Boolean indicating whether to show when monitors/groups last triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#show_last_triggered Powerpack#show_last_triggered}
	ShowLastTriggered interface{} `field:"optional" json:"showLastTriggered" yaml:"showLastTriggered"`
	// Whether to show the priorities column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#show_priority Powerpack#show_priority}
	ShowPriority interface{} `field:"optional" json:"showPriority" yaml:"showPriority"`
	// The method to sort the monitors.
	//
	// Valid values are `name`, `group`, `status`, `tags`, `triggered`, `group,asc`, `group,desc`, `name,asc`, `name,desc`, `status,asc`, `status,desc`, `tags,asc`, `tags,desc`, `triggered,asc`, `triggered,desc`, `priority,asc`, `priority,desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#sort Powerpack#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// The summary type to use. Valid values are `monitors`, `groups`, `combined`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#summary_type Powerpack#summary_type}
	SummaryType *string `field:"optional" json:"summaryType" yaml:"summaryType"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#title Powerpack#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#title_align Powerpack#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#title_size Powerpack#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

