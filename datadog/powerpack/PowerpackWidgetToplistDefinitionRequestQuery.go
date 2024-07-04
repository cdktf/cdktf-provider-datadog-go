// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetToplistDefinitionRequestQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#apm_dependency_stats_query Powerpack#apm_dependency_stats_query}
	ApmDependencyStatsQuery *PowerpackWidgetToplistDefinitionRequestQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#apm_resource_stats_query Powerpack#apm_resource_stats_query}
	ApmResourceStatsQuery *PowerpackWidgetToplistDefinitionRequestQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#cloud_cost_query Powerpack#cloud_cost_query}
	CloudCostQuery *PowerpackWidgetToplistDefinitionRequestQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#event_query Powerpack#event_query}
	EventQuery *PowerpackWidgetToplistDefinitionRequestQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#metric_query Powerpack#metric_query}
	MetricQuery *PowerpackWidgetToplistDefinitionRequestQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#process_query Powerpack#process_query}
	ProcessQuery *PowerpackWidgetToplistDefinitionRequestQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#slo_query Powerpack#slo_query}
	SloQuery *PowerpackWidgetToplistDefinitionRequestQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
}

