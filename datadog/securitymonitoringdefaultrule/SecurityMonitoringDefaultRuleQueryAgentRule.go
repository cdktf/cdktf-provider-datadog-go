// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringdefaultrule


type SecurityMonitoringDefaultRuleQueryAgentRule struct {
	// **Deprecated**. It won't be applied anymore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/security_monitoring_default_rule#agent_rule_id SecurityMonitoringDefaultRule#agent_rule_id}
	AgentRuleId *string `field:"required" json:"agentRuleId" yaml:"agentRuleId"`
	// **Deprecated**. It won't be applied anymore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/security_monitoring_default_rule#expression SecurityMonitoringDefaultRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

