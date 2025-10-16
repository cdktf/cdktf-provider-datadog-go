// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitornotificationrule


type MonitorNotificationRuleFilter struct {
	// All tags that target monitors must match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/monitor_notification_rule#tags MonitorNotificationRule#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
}

