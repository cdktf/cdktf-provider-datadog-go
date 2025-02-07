// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleOptionsImpossibleTravelOptions struct {
	// If true, signals are suppressed for the first 24 hours.
	//
	// During that time, Datadog learns the user's regular access locations. This can be helpful to reduce noise and infer VPN usage or credentialed API access. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/security_monitoring_rule#baseline_user_locations SecurityMonitoringRule#baseline_user_locations}
	BaselineUserLocations interface{} `field:"optional" json:"baselineUserLocations" yaml:"baselineUserLocations"`
}

