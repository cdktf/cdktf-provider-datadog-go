// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorSchedulingOptionsCustomSchedule struct {
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/monitor#recurrence Monitor#recurrence}
	Recurrence *MonitorSchedulingOptionsCustomScheduleRecurrence `field:"required" json:"recurrence" yaml:"recurrence"`
}

