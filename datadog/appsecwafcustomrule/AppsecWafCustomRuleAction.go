// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule


type AppsecWafCustomRuleAction struct {
	// Override the default action to take when the WAF custom rule would block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/appsec_waf_custom_rule#action AppsecWafCustomRule#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/appsec_waf_custom_rule#parameters AppsecWafCustomRule#parameters}
	Parameters *AppsecWafCustomRuleActionParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

