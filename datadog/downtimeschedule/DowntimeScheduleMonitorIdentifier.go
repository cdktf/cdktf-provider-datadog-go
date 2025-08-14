// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package downtimeschedule


type DowntimeScheduleMonitorIdentifier struct {
	// ID of the monitor to prevent notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/downtime_schedule#monitor_id DowntimeSchedule#monitor_id}
	MonitorId *float64 `field:"optional" json:"monitorId" yaml:"monitorId"`
	// A list of monitor tags.
	//
	// For example, tags that are applied directly to monitors, not tags that are used in monitor queries (which are filtered by the scope parameter), to which the downtime applies. The resulting downtime applies to monitors that match **all** provided monitor tags. Setting `monitor_tags` to `[*]` configures the downtime to mute all monitors for the given scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/downtime_schedule#monitor_tags DowntimeSchedule#monitor_tags}
	MonitorTags *[]*string `field:"optional" json:"monitorTags" yaml:"monitorTags"`
}

