// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetDistributionDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#apm_query Dashboard#apm_query}
	ApmQuery *DashboardWidgetDistributionDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// apm_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#apm_stats_query Dashboard#apm_stats_query}
	ApmStatsQuery *DashboardWidgetDistributionDefinitionRequestApmStatsQuery `field:"optional" json:"apmStatsQuery" yaml:"apmStatsQuery"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#log_query Dashboard#log_query}
	LogQuery *DashboardWidgetDistributionDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetDistributionDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#q Dashboard#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#rum_query Dashboard#rum_query}
	RumQuery *DashboardWidgetDistributionDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#security_query Dashboard#security_query}
	SecurityQuery *DashboardWidgetDistributionDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/dashboard#style Dashboard#style}
	Style *DashboardWidgetDistributionDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

