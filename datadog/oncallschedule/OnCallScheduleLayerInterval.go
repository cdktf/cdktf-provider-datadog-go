// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallschedule


type OnCallScheduleLayerInterval struct {
	// The number of full days in each rotation period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/on_call_schedule#days OnCallSchedule#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// For intervals that are not expressible in whole days, this will be added to `days`. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/on_call_schedule#seconds OnCallSchedule#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

