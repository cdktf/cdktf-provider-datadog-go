package dashboard


type DashboardWidgetToplistDefinitionRequestQueryEventQuerySearch struct {
	// The events search string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#query Dashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

