package monitor


type MonitorMonitorThresholdWindows struct {
	// Describes how long an anomalous metric must be normal before the alert recovers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#recovery_window Monitor#recovery_window}
	RecoveryWindow *string `field:"optional" json:"recoveryWindow" yaml:"recoveryWindow"`
	// Describes how long a metric must be anomalous before an alert triggers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#trigger_window Monitor#trigger_window}
	TriggerWindow *string `field:"optional" json:"triggerWindow" yaml:"triggerWindow"`
}

