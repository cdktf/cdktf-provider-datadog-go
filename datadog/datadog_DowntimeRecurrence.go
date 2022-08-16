// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DowntimeRecurrence struct {
	// One of `days`, `weeks`, `months`, `years`, or `rrule`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/downtime#type Downtime#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// How often to repeat as an integer.
	//
	// For example to repeat every 3 days, select a `type` of `days` and a `period` of `3`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/downtime#period Downtime#period}
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// The RRULE standard for defining recurring events.
	//
	// For example, to have a recurring event on the first day of each month, use `FREQ=MONTHLY;INTERVAL=1`. Most common rrule options from the iCalendar Spec are supported. Attributes specifying the duration in RRULE are not supported (for example, `DTSTART`, `DTEND`, `DURATION`). Only applicable when `type` is `rrule`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/downtime#rrule Downtime#rrule}
	Rrule *string `field:"optional" json:"rrule" yaml:"rrule"`
	// The date at which the recurrence should end as a POSIX timestamp. `until_occurrences` and `until_date` are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/downtime#until_date Downtime#until_date}
	UntilDate *float64 `field:"optional" json:"untilDate" yaml:"untilDate"`
	// How many times the downtime will be rescheduled. `until_occurrences` and `until_date` are mutually exclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/downtime#until_occurrences Downtime#until_occurrences}
	UntilOccurrences *float64 `field:"optional" json:"untilOccurrences" yaml:"untilOccurrences"`
	// A list of week days to repeat on.
	//
	// Choose from: `Mon`, `Tue`, `Wed`, `Thu`, `Fri`, `Sat` or `Sun`. Only applicable when `type` is `weeks`. First letter must be capitalized.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/downtime#week_days Downtime#week_days}
	WeekDays *[]*string `field:"optional" json:"weekDays" yaml:"weekDays"`
}

