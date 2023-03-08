package monitor


type MonitorSchedulingOptionsEvaluationWindow struct {
	// The time of the day at which a one day cumulative evaluation window starts.
	//
	// Must be defined in UTC time in `HH:mm` format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#day_starts Monitor#day_starts}
	DayStarts *string `field:"optional" json:"dayStarts" yaml:"dayStarts"`
	// The minute of the hour at which a one hour cumulative evaluation window starts.
	//
	// Must be between 0 and 59.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#hour_starts Monitor#hour_starts}
	HourStarts *float64 `field:"optional" json:"hourStarts" yaml:"hourStarts"`
	// The day of the month at which a one month cumulative evaluation window starts.
	//
	// Must be a value of 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/monitor#month_starts Monitor#month_starts}
	MonthStarts *float64 `field:"optional" json:"monthStarts" yaml:"monthStarts"`
}

