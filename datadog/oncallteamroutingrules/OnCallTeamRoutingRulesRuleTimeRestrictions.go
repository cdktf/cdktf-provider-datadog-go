// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleTimeRestrictions struct {
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_team_routing_rules#restriction OnCallTeamRoutingRules#restriction}
	Restriction interface{} `field:"optional" json:"restriction" yaml:"restriction"`
	// Specifies the time zone applicable to the restrictions, e.g. `America/New_York`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_team_routing_rules#time_zone OnCallTeamRoutingRules#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

