// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetLogStreamDefinition struct {
	// Stringified list of columns to use, for example: `["column1","column2","column3"]`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#columns Dashboard#columns}
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// An array of index names to query in the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#indexes Dashboard#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#live_span Dashboard#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The number of log lines to display. Valid values are `inline`, `expanded-md`, `expanded-lg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#message_display Dashboard#message_display}
	MessageDisplay *string `field:"optional" json:"messageDisplay" yaml:"messageDisplay"`
	// The query to use in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#query Dashboard#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// If the date column should be displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#show_date_column Dashboard#show_date_column}
	ShowDateColumn interface{} `field:"optional" json:"showDateColumn" yaml:"showDateColumn"`
	// If the message column should be displayed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#show_message_column Dashboard#show_message_column}
	ShowMessageColumn interface{} `field:"optional" json:"showMessageColumn" yaml:"showMessageColumn"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#sort Dashboard#sort}
	Sort *DashboardWidgetGroupDefinitionWidgetLogStreamDefinitionSort `field:"optional" json:"sort" yaml:"sort"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#title_align Dashboard#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#title_size Dashboard#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

