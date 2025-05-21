// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule


type AppsecWafCustomRuleConditionParametersInput struct {
	// Input from the request on which the condition should apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/appsec_waf_custom_rule#address AppsecWafCustomRule#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Specific path for the input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/appsec_waf_custom_rule#key_path AppsecWafCustomRule#key_path}
	KeyPath *[]*string `field:"optional" json:"keyPath" yaml:"keyPath"`
}

