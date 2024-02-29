// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorSchedulingOptions struct {
	// custom_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/monitor#custom_schedule Monitor#custom_schedule}
	CustomSchedule interface{} `field:"optional" json:"customSchedule" yaml:"customSchedule"`
	// evaluation_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/monitor#evaluation_window Monitor#evaluation_window}
	EvaluationWindow interface{} `field:"optional" json:"evaluationWindow" yaml:"evaluationWindow"`
}

