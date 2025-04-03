// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule


type AppsecWafCustomRuleConditionParameters struct {
	// Identifier of a list of data from the denylist. Can only be used as substitution from the list parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_custom_rule#data AppsecWafCustomRule#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
	// input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_custom_rule#input AppsecWafCustomRule#input}
	Input interface{} `field:"optional" json:"input" yaml:"input"`
	// List of value to use with the condition. Only used with the phrase_match, !phrase_match, exact_match and !exact_match operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_custom_rule#list AppsecWafCustomRule#list}
	List *[]*string `field:"optional" json:"list" yaml:"list"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_custom_rule#options AppsecWafCustomRule#options}
	Options *AppsecWafCustomRuleConditionParametersOptions `field:"optional" json:"options" yaml:"options"`
	// Regex to use with the condition. Only used with match_regex and !match_regex operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_custom_rule#regex AppsecWafCustomRule#regex}
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
	// Store the captured value in the specified tag name. Only used with the capture_data operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_custom_rule#value AppsecWafCustomRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

