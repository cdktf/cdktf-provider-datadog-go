// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetEventTimelineDefinition struct {
	// The query to use in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/powerpack#query Powerpack#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/powerpack#live_span Powerpack#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The execution method for multi-value filters, options: `and` or `or`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/powerpack#tags_execution Powerpack#tags_execution}
	TagsExecution *string `field:"optional" json:"tagsExecution" yaml:"tagsExecution"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/powerpack#title Powerpack#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/powerpack#title_align Powerpack#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/powerpack#title_size Powerpack#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

