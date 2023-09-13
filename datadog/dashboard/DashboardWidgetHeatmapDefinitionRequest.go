// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetHeatmapDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#apm_query Dashboard#apm_query}
	ApmQuery *DashboardWidgetHeatmapDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#log_query Dashboard#log_query}
	LogQuery *DashboardWidgetHeatmapDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetHeatmapDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#q Dashboard#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#rum_query Dashboard#rum_query}
	RumQuery *DashboardWidgetHeatmapDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#security_query Dashboard#security_query}
	SecurityQuery *DashboardWidgetHeatmapDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#style Dashboard#style}
	Style *DashboardWidgetHeatmapDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

