// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package downtimeschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DowntimeScheduleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The scope to which the downtime applies. Must follow the [common search syntax](https://docs.datadoghq.com/logs/explorer/search_syntax/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#scope DowntimeSchedule#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// The timezone in which to display the downtime's start and end times in Datadog applications.
	//
	// This is not used as an offset for scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#display_timezone DowntimeSchedule#display_timezone}
	DisplayTimezone *string `field:"optional" json:"displayTimezone" yaml:"displayTimezone"`
	// A message to include with notifications for this downtime.
	//
	// Email notifications can be sent to specific users by using the same `@username` notation as events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#message DowntimeSchedule#message}
	Message *string `field:"optional" json:"message" yaml:"message"`
	// monitor_identifier block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#monitor_identifier DowntimeSchedule#monitor_identifier}
	MonitorIdentifier *DowntimeScheduleMonitorIdentifier `field:"optional" json:"monitorIdentifier" yaml:"monitorIdentifier"`
	// If the first recovery notification during a downtime should be muted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#mute_first_recovery_notification DowntimeSchedule#mute_first_recovery_notification}
	MuteFirstRecoveryNotification interface{} `field:"optional" json:"muteFirstRecoveryNotification" yaml:"muteFirstRecoveryNotification"`
	// States that will trigger a monitor notification when the `notify_end_types` action occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#notify_end_states DowntimeSchedule#notify_end_states}
	NotifyEndStates *[]*string `field:"optional" json:"notifyEndStates" yaml:"notifyEndStates"`
	// Actions that will trigger a monitor notification if the downtime is in the `notify_end_types` state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#notify_end_types DowntimeSchedule#notify_end_types}
	NotifyEndTypes *[]*string `field:"optional" json:"notifyEndTypes" yaml:"notifyEndTypes"`
	// one_time_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#one_time_schedule DowntimeSchedule#one_time_schedule}
	OneTimeSchedule *DowntimeScheduleOneTimeSchedule `field:"optional" json:"oneTimeSchedule" yaml:"oneTimeSchedule"`
	// recurring_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/downtime_schedule#recurring_schedule DowntimeSchedule#recurring_schedule}
	RecurringSchedule *DowntimeScheduleRecurringSchedule `field:"optional" json:"recurringSchedule" yaml:"recurringSchedule"`
}

