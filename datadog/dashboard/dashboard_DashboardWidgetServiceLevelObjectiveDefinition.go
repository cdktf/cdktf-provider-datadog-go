package dashboard


type DashboardWidgetServiceLevelObjectiveDefinition struct {
	// The ID of the service level objective used by the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#slo_id Dashboard#slo_id}
	SloId *string `field:"required" json:"sloId" yaml:"sloId"`
	// A list of time windows to display in the widget.
	//
	// Valid values are `7d`, `30d`, `90d`, `week_to_date`, `previous_week`, `month_to_date`, `previous_month`, `global_time`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#time_windows Dashboard#time_windows}
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// The view mode for the widget. Valid values are `overall`, `component`, `both`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#view_mode Dashboard#view_mode}
	ViewMode *string `field:"required" json:"viewMode" yaml:"viewMode"`
	// The type of view to use when displaying the widget. Only `detail` is supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#view_type Dashboard#view_type}
	ViewType *string `field:"required" json:"viewType" yaml:"viewType"`
	// The global time target of the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#global_time_target Dashboard#global_time_target}
	GlobalTimeTarget *string `field:"optional" json:"globalTimeTarget" yaml:"globalTimeTarget"`
	// Whether to show the error budget or not.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#show_error_budget Dashboard#show_error_budget}
	ShowErrorBudget interface{} `field:"optional" json:"showErrorBudget" yaml:"showErrorBudget"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title_align Dashboard#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#title_size Dashboard#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}
