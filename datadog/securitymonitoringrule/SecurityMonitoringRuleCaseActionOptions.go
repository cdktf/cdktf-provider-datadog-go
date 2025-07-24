// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleCaseActionOptions struct {
	// Duration of the action in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/security_monitoring_rule#duration SecurityMonitoringRule#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
}

