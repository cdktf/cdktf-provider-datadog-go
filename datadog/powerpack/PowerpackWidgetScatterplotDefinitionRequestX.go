// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetScatterplotDefinitionRequestX struct {
	// Aggregator used for the request. Valid values are `avg`, `last`, `max`, `min`, `sum`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/powerpack#aggregator Powerpack#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/powerpack#apm_query Powerpack#apm_query}
	ApmQuery *PowerpackWidgetScatterplotDefinitionRequestXApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/powerpack#log_query Powerpack#log_query}
	LogQuery *PowerpackWidgetScatterplotDefinitionRequestXLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/powerpack#process_query Powerpack#process_query}
	ProcessQuery *PowerpackWidgetScatterplotDefinitionRequestXProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/powerpack#q Powerpack#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/powerpack#rum_query Powerpack#rum_query}
	RumQuery *PowerpackWidgetScatterplotDefinitionRequestXRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/powerpack#security_query Powerpack#security_query}
	SecurityQuery *PowerpackWidgetScatterplotDefinitionRequestXSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
}

