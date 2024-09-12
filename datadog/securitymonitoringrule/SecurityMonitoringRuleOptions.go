// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptions struct {
	// If true, signals in non-production environments have a lower severity than what is defined by the rule case, which can reduce noise.
	//
	// The decrement is applied when the environment tag of the signal starts with `staging`, `test`, or `dev`. Only available when the rule type is `log_detection`. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#decrease_criticality_based_on_env SecurityMonitoringRule#decrease_criticality_based_on_env}
	DecreaseCriticalityBasedOnEnv interface{} `field:"optional" json:"decreaseCriticalityBasedOnEnv" yaml:"decreaseCriticalityBasedOnEnv"`
	// The detection method. Valid values are `threshold`, `new_value`, `anomaly_detection`, `impossible_travel`, `hardcoded`, `third_party`, `anomaly_threshold`. Defaults to `"threshold"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#detection_method SecurityMonitoringRule#detection_method}
	DetectionMethod *string `field:"optional" json:"detectionMethod" yaml:"detectionMethod"`
	// A time window is specified to match when at least one of the cases matches true.
	//
	// This is a sliding window and evaluates in real time. Valid values are `0`, `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#evaluation_window SecurityMonitoringRule#evaluation_window}
	EvaluationWindow *float64 `field:"optional" json:"evaluationWindow" yaml:"evaluationWindow"`
	// impossible_travel_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#impossible_travel_options SecurityMonitoringRule#impossible_travel_options}
	ImpossibleTravelOptions *SecurityMonitoringRuleOptionsImpossibleTravelOptions `field:"optional" json:"impossibleTravelOptions" yaml:"impossibleTravelOptions"`
	// Once a signal is generated, the signal will remain “open” if a case is matched at least once within this keep alive window (in seconds).
	//
	// Valid values are `0`, `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `10800`, `21600`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#keep_alive SecurityMonitoringRule#keep_alive}
	KeepAlive *float64 `field:"optional" json:"keepAlive" yaml:"keepAlive"`
	// A signal will “close” regardless of the query being matched once the time exceeds the maximum duration (in seconds).
	//
	// This time is calculated from the first seen timestamp. Valid values are `0`, `60`, `300`, `600`, `900`, `1800`, `3600`, `7200`, `10800`, `21600`, `43200`, `86400`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#max_signal_duration SecurityMonitoringRule#max_signal_duration}
	MaxSignalDuration *float64 `field:"optional" json:"maxSignalDuration" yaml:"maxSignalDuration"`
	// new_value_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#new_value_options SecurityMonitoringRule#new_value_options}
	NewValueOptions *SecurityMonitoringRuleOptionsNewValueOptions `field:"optional" json:"newValueOptions" yaml:"newValueOptions"`
	// third_party_rule_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/security_monitoring_rule#third_party_rule_options SecurityMonitoringRule#third_party_rule_options}
	ThirdPartyRuleOptions *SecurityMonitoringRuleOptionsThirdPartyRuleOptions `field:"optional" json:"thirdPartyRuleOptions" yaml:"thirdPartyRuleOptions"`
}

