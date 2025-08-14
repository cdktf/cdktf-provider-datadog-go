// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#apm_query Dashboard#apm_query}
	ApmQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// audit_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#audit_query Dashboard#audit_query}
	AuditQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestAuditQuery `field:"optional" json:"auditQuery" yaml:"auditQuery"`
	// How to display the marker lines. Valid values are `area`, `bars`, `line`, `overlay`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#display_type Dashboard#display_type}
	DisplayType *string `field:"optional" json:"displayType" yaml:"displayType"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#formula Dashboard#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#log_query Dashboard#log_query}
	LogQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#metadata Dashboard#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// network_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#network_query Dashboard#network_query}
	NetworkQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestNetworkQuery `field:"optional" json:"networkQuery" yaml:"networkQuery"`
	// A Boolean indicating whether the request uses the right or left Y-Axis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#on_right_yaxis Dashboard#on_right_yaxis}
	OnRightYaxis interface{} `field:"optional" json:"onRightYaxis" yaml:"onRightYaxis"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#process_query Dashboard#process_query}
	ProcessQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#q Dashboard#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#query Dashboard#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#rum_query Dashboard#rum_query}
	RumQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#security_query Dashboard#security_query}
	SecurityQuery *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#style Dashboard#style}
	Style *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

