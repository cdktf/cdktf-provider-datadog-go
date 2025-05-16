// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetServiceLevelObjectiveDefinition struct {
	// The ID of the service level objective used by the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#slo_id Powerpack#slo_id}
	SloId *string `field:"required" json:"sloId" yaml:"sloId"`
	// A list of time windows to display in the widget.
	//
	// Valid values are `7d`, `30d`, `90d`, `week_to_date`, `previous_week`, `month_to_date`, `previous_month`, `global_time`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#time_windows Powerpack#time_windows}
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// The view mode for the widget. Valid values are `overall`, `component`, `both`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#view_mode Powerpack#view_mode}
	ViewMode *string `field:"required" json:"viewMode" yaml:"viewMode"`
	// The type of view to use when displaying the widget. Only `detail` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#view_type Powerpack#view_type}
	ViewType *string `field:"required" json:"viewType" yaml:"viewType"`
	// Additional filters applied to the SLO query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#additional_query_filters Powerpack#additional_query_filters}
	AdditionalQueryFilters *string `field:"optional" json:"additionalQueryFilters" yaml:"additionalQueryFilters"`
	// The global time target of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#global_time_target Powerpack#global_time_target}
	GlobalTimeTarget *string `field:"optional" json:"globalTimeTarget" yaml:"globalTimeTarget"`
	// Whether to show the error budget or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#show_error_budget Powerpack#show_error_budget}
	ShowErrorBudget interface{} `field:"optional" json:"showErrorBudget" yaml:"showErrorBudget"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#title Powerpack#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#title_align Powerpack#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/powerpack#title_size Powerpack#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

