// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentnotificationrule


type IncidentNotificationRuleConditions struct {
	// The incident field to evaluate. Common values include: state, severity, services, teams. Custom fields are also supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/incident_notification_rule#field IncidentNotificationRule#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// The value(s) to compare against. Multiple values are ORed together.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/incident_notification_rule#values IncidentNotificationRule#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

