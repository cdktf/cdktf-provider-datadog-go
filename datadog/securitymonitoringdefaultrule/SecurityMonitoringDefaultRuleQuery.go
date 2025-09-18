// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringdefaultrule


type SecurityMonitoringDefaultRuleQuery struct {
	// agent_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#agent_rule SecurityMonitoringDefaultRule#agent_rule}
	AgentRule interface{} `field:"optional" json:"agentRule" yaml:"agentRule"`
	// The aggregation type.
	//
	// For Signal Correlation rules, it must be event_count. Valid values are `count`, `cardinality`, `sum`, `max`, `new_value`, `geo_data`, `event_count`, `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#aggregation SecurityMonitoringDefaultRule#aggregation}
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Query extension to append to the logs query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#custom_query_extension SecurityMonitoringDefaultRule#custom_query_extension}
	CustomQueryExtension *string `field:"optional" json:"customQueryExtension" yaml:"customQueryExtension"`
	// Source of events. Valid values are `logs`, `audit`, `app_sec_spans`, `spans`, `security_runtime`, `network`, `events`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#data_source SecurityMonitoringDefaultRule#data_source}
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
	// Field for which the cardinality is measured. Sent as an array.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#distinct_fields SecurityMonitoringDefaultRule#distinct_fields}
	DistinctFields *[]*string `field:"optional" json:"distinctFields" yaml:"distinctFields"`
	// Fields to group by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#group_by_fields SecurityMonitoringDefaultRule#group_by_fields}
	GroupByFields *[]*string `field:"optional" json:"groupByFields" yaml:"groupByFields"`
	// When false, events without a group-by value are ignored by the rule.
	//
	// When true, events with missing group-by fields are processed with `N/A`, replacing the missing values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#has_optional_group_by_fields SecurityMonitoringDefaultRule#has_optional_group_by_fields}
	HasOptionalGroupByFields interface{} `field:"optional" json:"hasOptionalGroupByFields" yaml:"hasOptionalGroupByFields"`
	// The target field to aggregate over when using the `sum`, `max`, or `geo_data` aggregations.
	//
	// **Deprecated.** Configure `metrics` instead. This attribute will be removed in the next major version of the provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#metric SecurityMonitoringDefaultRule#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Group of target fields to aggregate over when using the `sum`, `max`, `geo_data`, or `new_value` aggregations.
	//
	// The `sum`, `max`, and `geo_data` aggregations only accept one value in this list, whereas the `new_value` aggregation accepts up to five values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#metrics SecurityMonitoringDefaultRule#metrics}
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Name of the query. Not compatible with `new_value` aggregations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#name SecurityMonitoringDefaultRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Query to run on logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/security_monitoring_default_rule#query SecurityMonitoringDefaultRule#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

