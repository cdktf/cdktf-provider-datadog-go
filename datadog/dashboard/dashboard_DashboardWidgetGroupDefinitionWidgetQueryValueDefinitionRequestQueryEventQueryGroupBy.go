package dashboard


type DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQueryGroupBy struct {
	// The event facet.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#facet Dashboard#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#sort Dashboard#sort}
	Sort *DashboardWidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryEventQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
}

