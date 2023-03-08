package monitor


type MonitorMonitorThresholds struct {
	// The monitor `CRITICAL` threshold. Must be a number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#critical Monitor#critical}
	Critical *string `field:"optional" json:"critical" yaml:"critical"`
	// The monitor `CRITICAL` recovery threshold. Must be a number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#critical_recovery Monitor#critical_recovery}
	CriticalRecovery *string `field:"optional" json:"criticalRecovery" yaml:"criticalRecovery"`
	// The monitor `OK` threshold. Only supported in monitor type `service check`. Must be a number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#ok Monitor#ok}
	Ok *string `field:"optional" json:"ok" yaml:"ok"`
	// The monitor `UNKNOWN` threshold. Only supported in monitor type `service check`. Must be a number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#unknown Monitor#unknown}
	Unknown *string `field:"optional" json:"unknown" yaml:"unknown"`
	// The monitor `WARNING` threshold. Must be a number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#warning Monitor#warning}
	Warning *string `field:"optional" json:"warning" yaml:"warning"`
	// The monitor `WARNING` recovery threshold. Must be a number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#warning_recovery Monitor#warning_recovery}
	WarningRecovery *string `field:"optional" json:"warningRecovery" yaml:"warningRecovery"`
}

