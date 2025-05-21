// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule


type AppsecWafCustomRuleCondition struct {
	// Operator to use for the WAF Condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/appsec_waf_custom_rule#operator AppsecWafCustomRule#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/appsec_waf_custom_rule#parameters AppsecWafCustomRule#parameters}
	Parameters *AppsecWafCustomRuleConditionParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

