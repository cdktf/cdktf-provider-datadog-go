package dashboard


type DashboardWidgetGroupDefinitionWidgetSloListDefinitionRequest struct {
	// query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#query Dashboard#query}
	Query *DashboardWidgetGroupDefinitionWidgetSloListDefinitionRequestQuery `field:"required" json:"query" yaml:"query"`
	// The request type for the SLO List request. Valid values are `slo_list`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#request_type Dashboard#request_type}
	RequestType *string `field:"required" json:"requestType" yaml:"requestType"`
}
