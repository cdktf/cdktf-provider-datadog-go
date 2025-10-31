// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/on_call_team_routing_rules#action OnCallTeamRoutingRules#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// ID of the policy to be applied when this routing rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/on_call_team_routing_rules#escalation_policy OnCallTeamRoutingRules#escalation_policy}
	EscalationPolicy *string `field:"optional" json:"escalationPolicy" yaml:"escalationPolicy"`
	// Defines the query or condition that triggers this routing rule. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/on_call_team_routing_rules#query OnCallTeamRoutingRules#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// time_restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/on_call_team_routing_rules#time_restrictions OnCallTeamRoutingRules#time_restrictions}
	TimeRestrictions *OnCallTeamRoutingRulesRuleTimeRestrictions `field:"optional" json:"timeRestrictions" yaml:"timeRestrictions"`
	// Defines the urgency for pages created via this rule.
	//
	// Only valid if `escalation_policy` is set. Valid values are `high`, `low`, `dynamic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/on_call_team_routing_rules#urgency OnCallTeamRoutingRules#urgency}
	Urgency *string `field:"optional" json:"urgency" yaml:"urgency"`
}

