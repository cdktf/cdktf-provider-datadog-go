// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleSignalQuery struct {
	// Rule ID of the signal to correlate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/security_monitoring_rule#rule_id SecurityMonitoringRule#rule_id}
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// The aggregation type.
	//
	// For Signal Correlation rules, it must be event_count. Valid values are `count`, `cardinality`, `sum`, `max`, `new_value`, `geo_data`, `event_count`, `none`. Defaults to `"event_count"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/security_monitoring_rule#aggregation SecurityMonitoringRule#aggregation}
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Fields to correlate by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/security_monitoring_rule#correlated_by_fields SecurityMonitoringRule#correlated_by_fields}
	CorrelatedByFields *[]*string `field:"optional" json:"correlatedByFields" yaml:"correlatedByFields"`
	// Index of the rule query used to retrieve the correlated field.
	//
	// An empty string applies correlation on the non-projected per query attributes of the rule. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/security_monitoring_rule#correlated_query_index SecurityMonitoringRule#correlated_query_index}
	CorrelatedQueryIndex *string `field:"optional" json:"correlatedQueryIndex" yaml:"correlatedQueryIndex"`
	// Default Rule ID of the signal to correlate. This value is READ-ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/security_monitoring_rule#default_rule_id SecurityMonitoringRule#default_rule_id}
	DefaultRuleId *string `field:"optional" json:"defaultRuleId" yaml:"defaultRuleId"`
	// Name of the query. Not compatible with `new_value` aggregations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

