package dashboard


type DashboardWidgetGroupDefinitionWidgetTreemapDefinitionRequestQueryEventQuerySearch struct {
	// The events search string.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#query Dashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

