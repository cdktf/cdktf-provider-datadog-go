// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetDistributionDefinitionRequest struct {
	// apm_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#apm_query Powerpack#apm_query}
	ApmQuery *PowerpackWidgetDistributionDefinitionRequestApmQuery `field:"optional" json:"apmQuery" yaml:"apmQuery"`
	// apm_stats_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#apm_stats_query Powerpack#apm_stats_query}
	ApmStatsQuery *PowerpackWidgetDistributionDefinitionRequestApmStatsQuery `field:"optional" json:"apmStatsQuery" yaml:"apmStatsQuery"`
	// log_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#log_query Powerpack#log_query}
	LogQuery *PowerpackWidgetDistributionDefinitionRequestLogQuery `field:"optional" json:"logQuery" yaml:"logQuery"`
	// process_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#process_query Powerpack#process_query}
	ProcessQuery *PowerpackWidgetDistributionDefinitionRequestProcessQuery `field:"optional" json:"processQuery" yaml:"processQuery"`
	// The metric query to use for this widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#q Powerpack#q}
	Q *string `field:"optional" json:"q" yaml:"q"`
	// rum_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#rum_query Powerpack#rum_query}
	RumQuery *PowerpackWidgetDistributionDefinitionRequestRumQuery `field:"optional" json:"rumQuery" yaml:"rumQuery"`
	// security_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#security_query Powerpack#security_query}
	SecurityQuery *PowerpackWidgetDistributionDefinitionRequestSecurityQuery `field:"optional" json:"securityQuery" yaml:"securityQuery"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#style Powerpack#style}
	Style *PowerpackWidgetDistributionDefinitionRequestStyle `field:"optional" json:"style" yaml:"style"`
}

