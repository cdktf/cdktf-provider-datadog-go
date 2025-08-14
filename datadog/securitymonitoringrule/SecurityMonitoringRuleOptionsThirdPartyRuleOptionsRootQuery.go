// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsThirdPartyRuleOptionsRootQuery struct {
	// Query to filter logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/security_monitoring_rule#query SecurityMonitoringRule#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// Fields to group by. If empty, each log triggers a signal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/security_monitoring_rule#group_by_fields SecurityMonitoringRule#group_by_fields}
	GroupByFields *[]*string `field:"optional" json:"groupByFields" yaml:"groupByFields"`
}

