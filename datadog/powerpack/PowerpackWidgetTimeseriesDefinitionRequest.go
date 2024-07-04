// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#apm_query Powerpack#apm_query}
	ApmQuery *PowerpackWidgetTimeseriesDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// audit_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#audit_query Powerpack#audit_query}
	AuditQuery *PowerpackWidgetTimeseriesDefinitionRequestAuditQuery `field:"optional" json:"auditQuery" yaml:"auditQuery"`
	// How to display the marker lines. Valid values are `area`, `bars`, `line`, `overlay`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#display_type Powerpack#display_type}
	DisplayType *string `field:"optional" json:"displayType" yaml:"displayType"`
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#formula Powerpack#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#log_query Powerpack#log_query}
	LogQuery *PowerpackWidgetTimeseriesDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#metadata Powerpack#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// network_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#network_query Powerpack#network_query}
	NetworkQuery *PowerpackWidgetTimeseriesDefinitionRequestNetworkQuery `field:"optional" json:"networkQuery" yaml:"networkQuery"`
	// A Boolean indicating whether the request uses the right or left Y-Axis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#on_right_yaxis Powerpack#on_right_yaxis}
	OnRightYaxis interface{} `field:"optional" json:"onRightYaxis" yaml:"onRightYaxis"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#process_query Powerpack#process_query}
	ProcessQuery *PowerpackWidgetTimeseriesDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#q Powerpack#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#query Powerpack#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#rum_query Powerpack#rum_query}
	RumQuery *PowerpackWidgetTimeseriesDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#security_query Powerpack#security_query}
	SecurityQuery *PowerpackWidgetTimeseriesDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/powerpack#style Powerpack#style}
	Style *PowerpackWidgetTimeseriesDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

