// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetTraceServiceDefinition struct {
	// APM environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#env Dashboard#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// APM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#service Dashboard#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// APM span name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#span_name Dashboard#span_name}
	SpanName *string `field:"required" json:"spanName" yaml:"spanName"`
	// The number of columns to display. Valid values are `one_column`, `two_column`, `three_column`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#display_format Dashboard#display_format}
	DisplayFormat *string `field:"optional" json:"displayFormat" yaml:"displayFormat"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `week_to_date`, `month_to_date`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#live_span Dashboard#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// Whether to show the latency breakdown or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#show_breakdown Dashboard#show_breakdown}
	ShowBreakdown interface{} `field:"optional" json:"showBreakdown" yaml:"showBreakdown"`
	// Whether to show the latency distribution or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#show_distribution Dashboard#show_distribution}
	ShowDistribution interface{} `field:"optional" json:"showDistribution" yaml:"showDistribution"`
	// Whether to show the error metrics or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#show_errors Dashboard#show_errors}
	ShowErrors interface{} `field:"optional" json:"showErrors" yaml:"showErrors"`
	// Whether to show the hits metrics or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#show_hits Dashboard#show_hits}
	ShowHits interface{} `field:"optional" json:"showHits" yaml:"showHits"`
	// Whether to show the latency metrics or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#show_latency Dashboard#show_latency}
	ShowLatency interface{} `field:"optional" json:"showLatency" yaml:"showLatency"`
	// Whether to show the resource list or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#show_resource_list Dashboard#show_resource_list}
	ShowResourceList interface{} `field:"optional" json:"showResourceList" yaml:"showResourceList"`
	// The size of the widget. Valid values are `small`, `medium`, `large`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#size_format Dashboard#size_format}
	SizeFormat *string `field:"optional" json:"sizeFormat" yaml:"sizeFormat"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#title_align Dashboard#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#title_size Dashboard#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

