// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetToplistDefinitionRequestQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/dashboard#apm_dependency_stats_query Dashboard#apm_dependency_stats_query}
	ApmDependencyStatsQuery *DashboardWidgetToplistDefinitionRequestQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/dashboard#apm_resource_stats_query Dashboard#apm_resource_stats_query}
	ApmResourceStatsQuery *DashboardWidgetToplistDefinitionRequestQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/dashboard#cloud_cost_query Dashboard#cloud_cost_query}
	CloudCostQuery *DashboardWidgetToplistDefinitionRequestQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/dashboard#event_query Dashboard#event_query}
	EventQuery *DashboardWidgetToplistDefinitionRequestQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/dashboard#metric_query Dashboard#metric_query}
	MetricQuery *DashboardWidgetToplistDefinitionRequestQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetToplistDefinitionRequestQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/dashboard#slo_query Dashboard#slo_query}
	SloQuery *DashboardWidgetToplistDefinitionRequestQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
}

