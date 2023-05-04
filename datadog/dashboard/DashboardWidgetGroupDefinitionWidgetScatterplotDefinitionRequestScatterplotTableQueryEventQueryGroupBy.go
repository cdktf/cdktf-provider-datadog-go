package dashboard


type DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryGroupBy struct {
	// The event facet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#facet Dashboard#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#sort Dashboard#sort}
	Sort *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
}

