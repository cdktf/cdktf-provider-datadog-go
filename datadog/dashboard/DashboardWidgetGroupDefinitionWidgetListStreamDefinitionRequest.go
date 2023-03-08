package dashboard


type DashboardWidgetGroupDefinitionWidgetListStreamDefinitionRequest struct {
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#columns Dashboard#columns}
	Columns interface{} `field:"required" json:"columns" yaml:"columns"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#query Dashboard#query}
	Query *DashboardWidgetGroupDefinitionWidgetListStreamDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// Widget response format. Valid values are `event_list`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#response_format Dashboard#response_format}
	ResponseFormat *string `field:"required" json:"responseFormat" yaml:"responseFormat"`
}

