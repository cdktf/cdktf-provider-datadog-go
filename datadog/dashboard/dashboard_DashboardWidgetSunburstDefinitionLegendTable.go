package dashboard


type DashboardWidgetSunburstDefinitionLegendTable struct {
	// The type of legend (table or none). Valid values are `table`, `none`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

