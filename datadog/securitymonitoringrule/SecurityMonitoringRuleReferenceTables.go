// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleReferenceTables struct {
	// Whether to include or exclude logs that match the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/security_monitoring_rule#check_presence SecurityMonitoringRule#check_presence}
	CheckPresence interface{} `field:"required" json:"checkPresence" yaml:"checkPresence"`
	// The name of the column in the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/security_monitoring_rule#column_name SecurityMonitoringRule#column_name}
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// The field in the log that should be matched against the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/security_monitoring_rule#log_field_path SecurityMonitoringRule#log_field_path}
	LogFieldPath *string `field:"required" json:"logFieldPath" yaml:"logFieldPath"`
	// The name of the query to filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/security_monitoring_rule#rule_query_name SecurityMonitoringRule#rule_query_name}
	RuleQueryName *string `field:"required" json:"ruleQueryName" yaml:"ruleQueryName"`
	// The name of the reference table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/security_monitoring_rule#table_name SecurityMonitoringRule#table_name}
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

