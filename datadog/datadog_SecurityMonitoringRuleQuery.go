// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type SecurityMonitoringRuleQuery struct {
	// Query to run on logs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#query SecurityMonitoringRule#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// agent_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#agent_rule SecurityMonitoringRule#agent_rule}
	AgentRule interface{} `field:"optional" json:"agentRule" yaml:"agentRule"`
	// The aggregation type. Valid values are `count`, `cardinality`, `sum`, `max`, `new_value`, `geo_data`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#aggregation SecurityMonitoringRule#aggregation}
	Aggregation *string `field:"optional" json:"aggregation" yaml:"aggregation"`
	// Field for which the cardinality is measured. Sent as an array.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#distinct_fields SecurityMonitoringRule#distinct_fields}
	DistinctFields *[]*string `field:"optional" json:"distinctFields" yaml:"distinctFields"`
	// Fields to group by.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#group_by_fields SecurityMonitoringRule#group_by_fields}
	GroupByFields *[]*string `field:"optional" json:"groupByFields" yaml:"groupByFields"`
	// The target field to aggregate over when using the `sum`, `max`, or `new_value` aggregations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#metric SecurityMonitoringRule#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Group of target fields to aggregate over when using the new value aggregations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#metrics SecurityMonitoringRule#metrics}
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// Name of the query. Not compatible with `new_value` aggregations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

