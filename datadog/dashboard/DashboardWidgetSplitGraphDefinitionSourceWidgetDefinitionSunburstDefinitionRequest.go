// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#apm_query Dashboard#apm_query}
	ApmQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// audit_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#audit_query Dashboard#audit_query}
	AuditQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestAuditQuery `field:"optional" json:"auditQuery" yaml:"auditQuery"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#formula Dashboard#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#log_query Dashboard#log_query}
	LogQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// network_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#network_query Dashboard#network_query}
	NetworkQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestNetworkQuery `field:"optional" json:"networkQuery" yaml:"networkQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#q Dashboard#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#query Dashboard#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#rum_query Dashboard#rum_query}
	RumQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#security_query Dashboard#security_query}
	SecurityQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#style Dashboard#style}
	Style *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

