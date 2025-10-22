// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallteamroutingrules


type OnCallTeamRoutingRulesRuleTimeRestrictionsRestriction struct {
	// The weekday when the restriction period ends. Valid values are `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/on_call_team_routing_rules#end_day OnCallTeamRoutingRules#end_day}
	EndDay *string `field:"optional" json:"endDay" yaml:"endDay"`
	// The time of day when the restriction ends (hh:mm:ss).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/on_call_team_routing_rules#end_time OnCallTeamRoutingRules#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The weekday when the restriction period starts. Valid values are `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/on_call_team_routing_rules#start_day OnCallTeamRoutingRules#start_day}
	StartDay *string `field:"optional" json:"startDay" yaml:"startDay"`
	// The time of day when the restriction begins (hh:mm:ss).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/on_call_team_routing_rules#start_time OnCallTeamRoutingRules#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

