// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallschedule


type OnCallScheduleLayerRestriction struct {
	// The weekday when the restriction period ends. Valid values are `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/on_call_schedule#end_day OnCallSchedule#end_day}
	EndDay *string `field:"required" json:"endDay" yaml:"endDay"`
	// The time of day when the restriction ends (hh:mm:ss).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/on_call_schedule#end_time OnCallSchedule#end_time}
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// The weekday when the restriction period starts. Valid values are `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/on_call_schedule#start_day OnCallSchedule#start_day}
	StartDay *string `field:"required" json:"startDay" yaml:"startDay"`
	// The time of day when the restriction begins (hh:mm:ss).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/on_call_schedule#start_time OnCallSchedule#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

