// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package downtimeschedule


type DowntimeScheduleRecurringScheduleRecurrence struct {
	// The length of the downtime.
	//
	// Must begin with an integer and end with one of 'm', 'h', d', or 'w'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/downtime_schedule#duration DowntimeSchedule#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// The `RRULE` standard for defining recurring events.
	//
	// For example, to have a recurring event on the first day of each month, set the type to `rrule` and set the `FREQ` to `MONTHLY` and `BYMONTHDAY` to `1`. Most common `rrule` options from the [iCalendar Spec](https://tools.ietf.org/html/rfc5545) are supported.  **Note**: Attributes specifying the duration in `RRULE` are not supported (for example, `DTSTART`, `DTEND`, `DURATION`). More examples available in this [downtime guide](https://docs.datadoghq.com/monitors/guide/suppress-alert-with-downtimes/?tab=api).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/downtime_schedule#rrule DowntimeSchedule#rrule}
	Rrule *string `field:"required" json:"rrule" yaml:"rrule"`
	// ISO-8601 Datetime to start the downtime.
	//
	// Must not include a UTC offset. If not provided, the downtime starts the moment it is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/downtime_schedule#start DowntimeSchedule#start}
	Start *string `field:"optional" json:"start" yaml:"start"`
}

