package securitymonitoringrule


type SecurityMonitoringRuleQueryAgentRule struct {
	// **Deprecated**. It won't be applied anymore.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#agent_rule_id SecurityMonitoringRule#agent_rule_id}
	AgentRuleId *string `field:"required" json:"agentRuleId" yaml:"agentRuleId"`
	// **Deprecated**. It won't be applied anymore.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/security_monitoring_rule#expression SecurityMonitoringRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

