// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsindex


type LogsIndexDailyLimitReset struct {
	// String in `HH:00` format representing the time of day the daily limit should be reset.
	//
	// The hours must be between 00 and 23 (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/logs_index#reset_time LogsIndex#reset_time}
	ResetTime *string `field:"required" json:"resetTime" yaml:"resetTime"`
	// String in `(-|+)HH:00` format representing the UTC offset to apply to the given reset time.
	//
	// The hours must be between -12 and +14 (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/logs_index#reset_utc_offset LogsIndex#reset_utc_offset}
	ResetUtcOffset *string `field:"required" json:"resetUtcOffset" yaml:"resetUtcOffset"`
}

