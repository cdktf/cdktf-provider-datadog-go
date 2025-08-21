// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package downtimeschedule


type DowntimeScheduleOneTimeSchedule struct {
	// ISO-8601 Datetime to end the downtime.
	//
	// Must include a UTC offset of zero. If not provided, the downtime never ends.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/downtime_schedule#end DowntimeSchedule#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// ISO-8601 Datetime to start the downtime.
	//
	// Must include a UTC offset of zero. If not provided, the downtime starts the moment it is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/downtime_schedule#start DowntimeSchedule#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
}

