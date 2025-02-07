// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorSchedulingOptionsCustomScheduleRecurrence struct {
	// Must be a valid `rrule`. See API docs for supported fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/monitor#rrule Monitor#rrule}
	Rrule *string `field:"required" json:"rrule" yaml:"rrule"`
	// 'tz database' format. Example: `America/New_York` or `UTC`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/monitor#timezone Monitor#timezone}
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
	// Time to start recurrence cycle. Similar to DTSTART. Expected format 'YYYY-MM-DDThh:mm:ss'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/monitor#start Monitor#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
}

