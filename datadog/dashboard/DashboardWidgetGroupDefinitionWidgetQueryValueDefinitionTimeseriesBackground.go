package dashboard


type DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionTimeseriesBackground struct {
	// Whether the Timeseries is made using an area or bars. Valid values are `bars`, `area`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#yaxis Dashboard#yaxis}
	Yaxis *DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionTimeseriesBackgroundYaxis `field:"optional" json:"yaxis" yaml:"yaxis"`
}

