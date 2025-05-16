// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule


type AppsecWafCustomRuleConditionParametersOptions struct {
	// Evaluate the value as case sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/appsec_waf_custom_rule#case_sensitive AppsecWafCustomRule#case_sensitive}
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Only evaluate this condition if the value has a minimum amount of characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/appsec_waf_custom_rule#min_length AppsecWafCustomRule#min_length}
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
}

