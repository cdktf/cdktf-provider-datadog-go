package dashboard


type DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#apm_dependency_stats_query Dashboard#apm_dependency_stats_query}
	ApmDependencyStatsQuery *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#apm_resource_stats_query Dashboard#apm_resource_stats_query}
	ApmResourceStatsQuery *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#event_query Dashboard#event_query}
	EventQuery *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#metric_query Dashboard#metric_query}
	MetricQuery *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
}

