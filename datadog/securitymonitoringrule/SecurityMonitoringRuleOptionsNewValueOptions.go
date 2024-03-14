// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsNewValueOptions struct {
	// The duration in days after which a learned value is forgotten.
	//
	// Valid values are `1`, `2`, `7`, `14`, `21`, `28`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#forget_after SecurityMonitoringRule#forget_after}
	ForgetAfter *float64 `field:"required" json:"forgetAfter" yaml:"forgetAfter"`
	// The duration in days during which values are learned, and after which signals will be generated for values that weren't learned.
	//
	// If set to 0, a signal will be generated for all new values after the first value is learned. Valid values are `0`, `1`, `7`. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#learning_duration SecurityMonitoringRule#learning_duration}
	LearningDuration *float64 `field:"optional" json:"learningDuration" yaml:"learningDuration"`
	// The learning method used to determine when signals should be generated for values that weren't learned.
	//
	// Valid values are `duration`, `threshold`. Defaults to `"duration"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#learning_method SecurityMonitoringRule#learning_method}
	LearningMethod *string `field:"optional" json:"learningMethod" yaml:"learningMethod"`
	// A number of occurrences after which signals are generated for values that weren't learned.
	//
	// Valid values are `0`, `1`. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#learning_threshold SecurityMonitoringRule#learning_threshold}
	LearningThreshold *float64 `field:"optional" json:"learningThreshold" yaml:"learningThreshold"`
}

