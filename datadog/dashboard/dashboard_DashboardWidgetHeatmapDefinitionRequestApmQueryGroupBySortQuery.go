package dashboard


type DashboardWidgetHeatmapDefinitionRequestApmQueryGroupBySortQuery struct {
	// The aggregation method.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#aggregation Dashboard#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#order Dashboard#order}
	Order *string `field:"required" json:"order" yaml:"order"`
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#facet Dashboard#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
}

