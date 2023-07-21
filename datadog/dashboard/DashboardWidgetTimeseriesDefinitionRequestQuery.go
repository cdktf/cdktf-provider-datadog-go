package dashboard


type DashboardWidgetTimeseriesDefinitionRequestQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#apm_dependency_stats_query Dashboard#apm_dependency_stats_query}
	ApmDependencyStatsQuery *DashboardWidgetTimeseriesDefinitionRequestQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#apm_resource_stats_query Dashboard#apm_resource_stats_query}
	ApmResourceStatsQuery *DashboardWidgetTimeseriesDefinitionRequestQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#event_query Dashboard#event_query}
	EventQuery *DashboardWidgetTimeseriesDefinitionRequestQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#metric_query Dashboard#metric_query}
	MetricQuery *DashboardWidgetTimeseriesDefinitionRequestQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetTimeseriesDefinitionRequestQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/dashboard#slo_query Dashboard#slo_query}
	SloQuery *DashboardWidgetTimeseriesDefinitionRequestQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
}

