// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestOptionsListMonitorOptions struct {
	// A message to include with a re-notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/synthetics_test#escalation_message SyntheticsTest#escalation_message}
	EscalationMessage *string `field:"optional" json:"escalationMessage" yaml:"escalationMessage"`
	// The name of the preset for the notification for the monitor. Valid values are `show_all`, `hide_all`, `hide_query`, `hide_handles`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/synthetics_test#notification_preset_name SyntheticsTest#notification_preset_name}
	NotificationPresetName *string `field:"optional" json:"notificationPresetName" yaml:"notificationPresetName"`
	// Specify a renotification frequency in minutes.
	//
	// Values available by default are `0`, `10`, `20`, `30`, `40`, `50`, `60`, `90`, `120`, `180`, `240`, `300`, `360`, `720`, `1440`. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/synthetics_test#renotify_interval SyntheticsTest#renotify_interval}
	RenotifyInterval *float64 `field:"optional" json:"renotifyInterval" yaml:"renotifyInterval"`
	// The number of times a monitor renotifies. It can only be set if `renotify_interval` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/synthetics_test#renotify_occurrences SyntheticsTest#renotify_occurrences}
	RenotifyOccurrences *float64 `field:"optional" json:"renotifyOccurrences" yaml:"renotifyOccurrences"`
}

