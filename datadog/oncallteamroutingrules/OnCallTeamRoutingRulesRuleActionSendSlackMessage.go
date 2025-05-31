// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleActionSendSlackMessage struct {
	// Slack channel ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_team_routing_rules#channel OnCallTeamRoutingRules#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// Slack workspace ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_team_routing_rules#workspace OnCallTeamRoutingRules#workspace}
	Workspace *string `field:"optional" json:"workspace" yaml:"workspace"`
}

