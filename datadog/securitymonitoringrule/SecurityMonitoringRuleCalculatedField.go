// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule


type SecurityMonitoringRuleCalculatedField struct {
	// Expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/security_monitoring_rule#expression SecurityMonitoringRule#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Field name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

