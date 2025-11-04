// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleSchedulingOptions struct {
	// Schedule for the rule queries, written in RRULE syntax. See [RFC](https://icalendar.org/iCalendar-RFC-5545/3-8-5-3-recurrence-rule.html) for syntax reference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/security_monitoring_rule#rrule SecurityMonitoringRule#rrule}
	Rrule *string `field:"required" json:"rrule" yaml:"rrule"`
	// Start date for the schedule, in ISO 8601 format without timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/security_monitoring_rule#start SecurityMonitoringRule#start}
	Start *string `field:"required" json:"start" yaml:"start"`
	// Time zone of the start date, in the [tz database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/security_monitoring_rule#timezone SecurityMonitoringRule#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
}

