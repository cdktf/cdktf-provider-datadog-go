package dashboard


type DashboardWidgetGeomapDefinitionRequestFormulaLimit struct {
	// The number of results to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#count Dashboard#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The direction of the sort. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#order Dashboard#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

