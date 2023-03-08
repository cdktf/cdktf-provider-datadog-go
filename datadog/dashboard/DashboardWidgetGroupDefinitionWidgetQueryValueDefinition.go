package dashboard


type DashboardWidgetGroupDefinitionWidgetQueryValueDefinition struct {
	// A Boolean indicating whether to automatically scale the tile.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#autoscale Dashboard#autoscale}
	Autoscale interface{} `field:"optional" json:"autoscale" yaml:"autoscale"`
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#custom_link Dashboard#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// The unit for the value displayed in the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#custom_unit Dashboard#custom_unit}
	CustomUnit *string `field:"optional" json:"customUnit" yaml:"customUnit"`
	// The timeframe to use when displaying the widget.
	//
	// Valid values are `1m`, `5m`, `10m`, `15m`, `30m`, `1h`, `4h`, `1d`, `2d`, `1w`, `1mo`, `3mo`, `6mo`, `1y`, `alert`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#live_span Dashboard#live_span}
	LiveSpan *string `field:"optional" json:"liveSpan" yaml:"liveSpan"`
	// The precision to use when displaying the tile.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#precision Dashboard#precision}
	Precision *float64 `field:"optional" json:"precision" yaml:"precision"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#request Dashboard#request}
	Request interface{} `field:"optional" json:"request" yaml:"request"`
	// The alignment of the widget's text. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#text_align Dashboard#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
	// timeseries_background block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#timeseries_background Dashboard#timeseries_background}
	TimeseriesBackground *DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionTimeseriesBackground `field:"optional" json:"timeseriesBackground" yaml:"timeseriesBackground"`
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

