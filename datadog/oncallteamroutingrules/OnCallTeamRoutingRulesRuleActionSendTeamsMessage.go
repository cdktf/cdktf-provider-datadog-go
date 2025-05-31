// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleActionSendTeamsMessage struct {
	// Teams channel ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_team_routing_rules#channel OnCallTeamRoutingRules#channel}
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// Teams team ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_team_routing_rules#team OnCallTeamRoutingRules#team}
	Team *string `field:"optional" json:"team" yaml:"team"`
	// Teams tenant ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_team_routing_rules#tenant OnCallTeamRoutingRules#tenant}
	Tenant *string `field:"optional" json:"tenant" yaml:"tenant"`
}

