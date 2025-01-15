// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudconfigurationrule


type CloudConfigurationRuleFilter struct {
	// The type of filtering action. Valid values are `require`, `suppress`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/cloud_configuration_rule#action CloudConfigurationRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// Query for selecting logs to apply the filtering action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.0/docs/resources/cloud_configuration_rule#query CloudConfigurationRule#query}
	Query *string `field:"required" json:"query" yaml:"query"`
}

