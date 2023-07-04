package dashboard


type DashboardWidgetSloListDefinitionRequestQuery struct {
	// Widget query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#query_string Dashboard#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Maximum number of results to display in the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#sort Dashboard#sort}
	Sort *DashboardWidgetSloListDefinitionRequestQuerySort `field:"optional" json:"sort" yaml:"sort"`
}

