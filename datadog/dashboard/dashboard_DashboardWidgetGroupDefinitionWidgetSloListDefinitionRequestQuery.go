package dashboard


type DashboardWidgetGroupDefinitionWidgetSloListDefinitionRequestQuery struct {
	// Widget query.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#query_string Dashboard#query_string}
	QueryString *string `field:"required" json:"queryString" yaml:"queryString"`
	// Maximum number of results to display in the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
}

