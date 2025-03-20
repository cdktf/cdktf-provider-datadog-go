// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package downtimeschedule


type DowntimeScheduleRecurringSchedule struct {
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/downtime_schedule#recurrence DowntimeSchedule#recurrence}
	Recurrence interface{} `field:"optional" json:"recurrence" yaml:"recurrence"`
	// The timezone in which to schedule the downtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/downtime_schedule#timezone DowntimeSchedule#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

