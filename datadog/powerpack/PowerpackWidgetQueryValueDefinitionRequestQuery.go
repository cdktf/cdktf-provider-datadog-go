// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryValueDefinitionRequestQuery struct {
	// apm_dependency_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#apm_dependency_stats_query Powerpack#apm_dependency_stats_query}
	ApmDependencyStatsQuery *PowerpackWidgetQueryValueDefinitionRequestQueryApmDependencyStatsQuery `field:"optional" json:"apmDependencyStatsQuery" yaml:"apmDependencyStatsQuery"`
	// apm_resource_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#apm_resource_stats_query Powerpack#apm_resource_stats_query}
	ApmResourceStatsQuery *PowerpackWidgetQueryValueDefinitionRequestQueryApmResourceStatsQuery `field:"optional" json:"apmResourceStatsQuery" yaml:"apmResourceStatsQuery"`
	// cloud_cost_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#cloud_cost_query Powerpack#cloud_cost_query}
	CloudCostQuery *PowerpackWidgetQueryValueDefinitionRequestQueryCloudCostQuery `field:"optional" json:"cloudCostQuery" yaml:"cloudCostQuery"`
	// event_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#event_query Powerpack#event_query}
	EventQuery *PowerpackWidgetQueryValueDefinitionRequestQueryEventQuery `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#metric_query Powerpack#metric_query}
	MetricQuery *PowerpackWidgetQueryValueDefinitionRequestQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#process_query Powerpack#process_query}
	ProcessQuery *PowerpackWidgetQueryValueDefinitionRequestQueryProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// slo_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#slo_query Powerpack#slo_query}
	SloQuery *PowerpackWidgetQueryValueDefinitionRequestQuerySloQuery `field:"optional" json:"sloQuery" yaml:"sloQuery"`
}

