// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleAction struct {
	// send_slack_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/on_call_team_routing_rules#send_slack_message OnCallTeamRoutingRules#send_slack_message}
	SendSlackMessage *OnCallTeamRoutingRulesRuleActionSendSlackMessage `field:"optional" json:"sendSlackMessage" yaml:"sendSlackMessage"`
	// send_teams_message block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/on_call_team_routing_rules#send_teams_message OnCallTeamRoutingRules#send_teams_message}
	SendTeamsMessage *OnCallTeamRoutingRulesRuleActionSendTeamsMessage `field:"optional" json:"sendTeamsMessage" yaml:"sendTeamsMessage"`
}

