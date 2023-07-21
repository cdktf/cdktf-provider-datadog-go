package monitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorConfig struct {
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
	// A message to include with notifications for this monitor.
	//
	// Email notifications can be sent to specific users by using the same `@username` notation as events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#message Monitor#message}
	Message *string `field:"required" json:"message" yaml:"message"`
	// Name of Datadog monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#name Monitor#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The monitor query to notify on.
	//
	// Note this is not the same query you see in the UI and the syntax is different depending on the monitor type, please see the [API Reference](https://docs.datadoghq.com/api/v1/monitors/#create-a-monitor) for details. `terraform plan` will validate query contents unless `validate` is set to `false`.
	//
	// *Note:** APM latency data is now available as Distribution Metrics. Existing monitors have been migrated automatically but all terraformed monitors can still use the existing metrics. We strongly recommend updating monitor definitions to query the new metrics. To learn more, or to see examples of how to update your terraform definitions to utilize the new distribution metrics, see the [detailed doc](https://docs.datadoghq.com/tracing/guide/ddsketch_trace_metrics/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#query Monitor#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The type of the monitor.
	//
	// The mapping from these types to the types found in the Datadog Web UI can be found in the Datadog API [documentation page](https://docs.datadoghq.com/api/v1/monitors/#create-a-monitor). Note: The monitor type cannot be changed after a monitor is created. Valid values are `composite`, `event alert`, `log alert`, `metric alert`, `process alert`, `query alert`, `rum alert`, `service check`, `synthetics alert`, `trace-analytics alert`, `slo alert`, `event-v2 alert`, `audit alert`, `ci-pipelines alert`, `ci-tests alert`, `error-tracking alert`, `database-monitoring alert`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#type Monitor#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A boolean indicating whether or not to include a list of log values which triggered the alert.
	//
	// This is only used by log monitors. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#enable_logs_sample Monitor#enable_logs_sample}
	EnableLogsSample interface{} `field:"optional" json:"enableLogsSample" yaml:"enableLogsSample"`
	// A message to include with a re-notification. Supports the `@username` notification allowed elsewhere.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#escalation_message Monitor#escalation_message}
	EscalationMessage *string `field:"optional" json:"escalationMessage" yaml:"escalationMessage"`
	// (Only applies to metric alert) Time (in seconds) to delay evaluation, as a non-negative integer.
	//
	// For example, if the value is set to `300` (5min), the `timeframe` is set to `last_5m` and the time is 7:00, the monitor will evaluate data from 6:50 to 6:55. This is useful for AWS CloudWatch and other backfilled metrics to ensure the monitor will always have data during evaluation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#evaluation_delay Monitor#evaluation_delay}
	EvaluationDelay *float64 `field:"optional" json:"evaluationDelay" yaml:"evaluationDelay"`
	// A boolean indicating whether this monitor can be deleted even if it’s referenced by other resources (e.g. SLO, composite monitor).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#force_delete Monitor#force_delete}
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Whether or not to trigger one alert if any source breaches a threshold.
	//
	// This is only used by log monitors. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#groupby_simple_monitor Monitor#groupby_simple_monitor}
	GroupbySimpleMonitor interface{} `field:"optional" json:"groupbySimpleMonitor" yaml:"groupbySimpleMonitor"`
	// The time span after which groups with missing data are dropped from the monitor state.
	//
	// The minimum value is one hour, and the maximum value is 72 hours. Example values are: 60m, 1h, and 2d. This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#group_retention_duration Monitor#group_retention_duration}
	GroupRetentionDuration *string `field:"optional" json:"groupRetentionDuration" yaml:"groupRetentionDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#id Monitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A boolean indicating whether notifications from this monitor automatically insert its triggering tags into the title. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#include_tags Monitor#include_tags}
	IncludeTags interface{} `field:"optional" json:"includeTags" yaml:"includeTags"`
	// A boolean indicating whether changes to this monitor should be restricted to the creator or admins.
	//
	// Defaults to `false`. **Deprecated.** Use `restricted_roles`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#locked Monitor#locked}
	Locked interface{} `field:"optional" json:"locked" yaml:"locked"`
	// monitor_thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#monitor_thresholds Monitor#monitor_thresholds}
	MonitorThresholds *MonitorMonitorThresholds `field:"optional" json:"monitorThresholds" yaml:"monitorThresholds"`
	// monitor_threshold_windows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#monitor_threshold_windows Monitor#monitor_threshold_windows}
	MonitorThresholdWindows *MonitorMonitorThresholdWindows `field:"optional" json:"monitorThresholdWindows" yaml:"monitorThresholdWindows"`
	// The time (in seconds) to skip evaluations for new groups.
	//
	// `new_group_delay` overrides `new_host_delay` if it is set to a nonzero value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#new_group_delay Monitor#new_group_delay}
	NewGroupDelay *float64 `field:"optional" json:"newGroupDelay" yaml:"newGroupDelay"`
	// **Deprecated**.
	//
	// See `new_group_delay`. Time (in seconds) to allow a host to boot and applications to fully start before starting the evaluation of monitor results. Should be a non-negative integer. This value is ignored for simple monitors and monitors not grouped by host. Defaults to `300`. The only case when this should be used is to override the default and set `new_host_delay` to zero for monitors grouped by host. **Deprecated.** Use `new_group_delay` except when setting `new_host_delay` to zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#new_host_delay Monitor#new_host_delay}
	NewHostDelay *float64 `field:"optional" json:"newHostDelay" yaml:"newHostDelay"`
	// The number of minutes before a monitor will notify when data stops reporting. Provider defaults to 10 minutes.
	//
	// We recommend at least 2x the monitor timeframe for metric alerts or 2 minutes for service checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#no_data_timeframe Monitor#no_data_timeframe}
	NoDataTimeframe *float64 `field:"optional" json:"noDataTimeframe" yaml:"noDataTimeframe"`
	// Toggles the display of additional content sent in the monitor notification. Valid values are `show_all`, `hide_query`, `hide_handles`, `hide_all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#notification_preset_name Monitor#notification_preset_name}
	NotificationPresetName *string `field:"optional" json:"notificationPresetName" yaml:"notificationPresetName"`
	// A boolean indicating whether tagged users will be notified on changes to this monitor. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#notify_audit Monitor#notify_audit}
	NotifyAudit interface{} `field:"optional" json:"notifyAudit" yaml:"notifyAudit"`
	// Controls what granularity a monitor alerts on.
	//
	// Only available for monitors with groupings. For instance, a monitor grouped by `cluster`, `namespace`, and `pod` can be configured to only notify on each new `cluster` violating the alert conditions by setting `notify_by` to `['cluster']`. Tags mentioned in `notify_by` must be a subset of the grouping tags in the query. For example, a query grouped by `cluster` and `namespace` cannot notify on `region`. Setting `notify_by` to `[*]` configures the monitor to notify as a simple-alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#notify_by Monitor#notify_by}
	NotifyBy *[]*string `field:"optional" json:"notifyBy" yaml:"notifyBy"`
	// A boolean indicating whether this monitor will notify when data stops reporting. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#notify_no_data Monitor#notify_no_data}
	NotifyNoData interface{} `field:"optional" json:"notifyNoData" yaml:"notifyNoData"`
	// Controls how groups or monitors are treated if an evaluation does not return any data points.
	//
	// The default option results in different behavior depending on the monitor query type. For monitors using `Count` queries, an empty monitor evaluation is treated as 0 and is compared to the threshold conditions. For monitors using any query type other than `Count`, for example `Gauge`, `Measure`, or `Rate`, the monitor shows the last known status. This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors. Valid values are: `show_no_data`, `show_and_notify_no_data`, `resolve`, and `default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#on_missing_data Monitor#on_missing_data}
	OnMissingData *string `field:"optional" json:"onMissingData" yaml:"onMissingData"`
	// Integer from 1 (high) to 5 (low) indicating alert severity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#priority Monitor#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The number of minutes after the last notification before a monitor will re-notify on the current status.
	//
	// It will only re-notify if it's not resolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#renotify_interval Monitor#renotify_interval}
	RenotifyInterval *float64 `field:"optional" json:"renotifyInterval" yaml:"renotifyInterval"`
	// The number of re-notification messages that should be sent on the current status.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#renotify_occurrences Monitor#renotify_occurrences}
	RenotifyOccurrences *float64 `field:"optional" json:"renotifyOccurrences" yaml:"renotifyOccurrences"`
	// The types of statuses for which re-notification messages should be sent. Valid values are `alert`, `warn`, `no data`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#renotify_statuses Monitor#renotify_statuses}
	RenotifyStatuses *[]*string `field:"optional" json:"renotifyStatuses" yaml:"renotifyStatuses"`
	// A boolean indicating whether this monitor needs a full window of data before it's evaluated.
	//
	// We highly recommend you set this to `false` for sparse metrics, otherwise some evaluations will be skipped. Default: `true` for `on average`, `at all times` and `in total` aggregation. `false` otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#require_full_window Monitor#require_full_window}
	RequireFullWindow interface{} `field:"optional" json:"requireFullWindow" yaml:"requireFullWindow"`
	// A list of unique role identifiers to define which roles are allowed to edit the monitor.
	//
	// Editing a monitor includes any updates to the monitor configuration, monitor deletion, and muting of the monitor for any amount of time. Roles unique identifiers can be pulled from the [Roles API](https://docs.datadoghq.com/api/latest/roles/#list-roles) in the `data.id` field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#restricted_roles Monitor#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// scheduling_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#scheduling_options Monitor#scheduling_options}
	SchedulingOptions interface{} `field:"optional" json:"schedulingOptions" yaml:"schedulingOptions"`
	// A list of tags to associate with your monitor.
	//
	// This can help you categorize and filter monitors in the manage monitors page of the UI. Note: it's not currently possible to filter by these tags when querying via the API
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#tags Monitor#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The number of hours of the monitor not reporting data before it automatically resolves from a triggered state.
	//
	// The minimum allowed value is 0 hours. The maximum allowed value is 24 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#timeout_h Monitor#timeout_h}
	TimeoutH *float64 `field:"optional" json:"timeoutH" yaml:"timeoutH"`
	// If set to `false`, skip the validation call done during plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#validate Monitor#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
	// variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.28.0/docs/resources/monitor#variables Monitor#variables}
	Variables *MonitorVariables `field:"optional" json:"variables" yaml:"variables"`
}

