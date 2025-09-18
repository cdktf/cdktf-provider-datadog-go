// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidentnotificationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IncidentNotificationRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The notification handles (targets) for this rule.
	//
	// Examples:.
	Handles *[]*string `field:"required" json:"handles" yaml:"handles"`
	// The ID of the incident type this notification rule is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/incident_notification_rule#incident_type IncidentNotificationRule#incident_type}
	IncidentType *string `field:"required" json:"incidentType" yaml:"incidentType"`
	// The trigger event for this notification rule. Valid values are: incident_created_trigger, incident_saved_trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/incident_notification_rule#trigger IncidentNotificationRule#trigger}
	Trigger *string `field:"required" json:"trigger" yaml:"trigger"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/incident_notification_rule#conditions IncidentNotificationRule#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Whether the notification rule is enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/incident_notification_rule#enabled IncidentNotificationRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID of the notification template to use for this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/incident_notification_rule#notification_template IncidentNotificationRule#notification_template}
	NotificationTemplate *string `field:"optional" json:"notificationTemplate" yaml:"notificationTemplate"`
	// List of incident fields that trigger re-notification when changed.
	//
	// Valid values are: status, severity, customer_impact, title, description, detected, root_cause, services, state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/incident_notification_rule#renotify_on IncidentNotificationRule#renotify_on}
	RenotifyOn *[]*string `field:"optional" json:"renotifyOn" yaml:"renotifyOn"`
	// The visibility of the notification rule. Valid values are: all, organization, private. Defaults to organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/incident_notification_rule#visibility IncidentNotificationRule#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

