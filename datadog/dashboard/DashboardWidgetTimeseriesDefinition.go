package dashboard


type DashboardWidgetTimeseriesDefinition struct {
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#custom_link Dashboard#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// event block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#event Dashboard#event}
	Event interface{} `field:"optional" json:"event" yaml:"event"`
	// A list of columns to display in the legend. Valid values are `value`, `avg`, `sum`, `min`, `max`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#legend_columns Dashboard#legend_columns}
	LegendColumns *[]*string `field:"optional" json:"legendColumns" yaml:"legendColumns"`
	// The layout of the legend displayed in the widget. Valid values are `auto`, `horizontal`, `vertical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#legend_layout Dashboard#legend_layout}
	LegendLayout *string `field:"optional" json:"legendLayout" yaml:"legendLayout"`
	// The size of the legend displayed in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#legend_size Dashboard#legend_size}
	LegendSize *string `field:"optional" json:"legendSize" yaml:"legendSize"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#live_span Dashboard#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// marker block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#marker Dashboard#marker}
	Marker interface{} `field:"optional" json:"marker" yaml:"marker"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#request Dashboard#request}
	Request interface{} `field:"optional" json:"request" yaml:"request"`
	// right_yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#right_yaxis Dashboard#right_yaxis}
	RightYaxis *DashboardWidgetTimeseriesDefinitionRightYaxis `field:"optional" json:"rightYaxis" yaml:"rightYaxis"`
	// Whether or not to show the legend on this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#show_legend Dashboard#show_legend}
	ShowLegend interface{} `field:"optional" json:"showLegend" yaml:"showLegend"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#title Dashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#title_align Dashboard#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#title_size Dashboard#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
	// yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#yaxis Dashboard#yaxis}
	Yaxis *DashboardWidgetTimeseriesDefinitionYaxis `field:"optional" json:"yaxis" yaml:"yaxis"`
}

