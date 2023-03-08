package dashboard


type DashboardWidgetGeomapDefinitionRequestQueryMetricQuery struct {
	// The name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metrics query definition.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#query Dashboard#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The aggregation methods available for metrics queries. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#aggregator Dashboard#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// The data source for metrics queries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
}

