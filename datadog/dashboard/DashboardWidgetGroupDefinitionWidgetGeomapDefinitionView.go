package dashboard


type DashboardWidgetGroupDefinitionWidgetGeomapDefinitionView struct {
	// The two-letter ISO code of a country to focus the map on (or `WORLD`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.25.0/docs/resources/dashboard#focus Dashboard#focus}
	Focus *string `field:"required" json:"focus" yaml:"focus"`
}

