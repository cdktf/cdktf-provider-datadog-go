package dashboard


type DashboardWidgetGroupDefinitionWidgetToplistDefinitionRequestApmQueryGroupBy struct {
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#facet Dashboard#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// The maximum number of items in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#sort_query Dashboard#sort_query}
	SortQuery *DashboardWidgetGroupDefinitionWidgetToplistDefinitionRequestApmQueryGroupBySortQuery `field:"optional" json:"sortQuery" yaml:"sortQuery"`
}

