package dashboard


type DashboardWidgetGroupDefinitionWidgetLogStreamDefinitionSort struct {
	// The facet path for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#column Dashboard#column}
	Column *string `field:"required" json:"column" yaml:"column"`
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#order Dashboard#order}
	Order *string `field:"required" json:"order" yaml:"order"`
}

