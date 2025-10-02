// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule


type AppsecWafCustomRuleScope struct {
	// The environment scope for the WAF custom rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#env AppsecWafCustomRule#env}
	Env *string `field:"optional" json:"env" yaml:"env"`
	// The service scope for the WAF custom rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#service AppsecWafCustomRule#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

