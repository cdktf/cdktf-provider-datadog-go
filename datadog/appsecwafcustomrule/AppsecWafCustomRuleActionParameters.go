// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule


type AppsecWafCustomRuleActionParameters struct {
	// The location to redirect to when the WAF custom rule triggers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/appsec_waf_custom_rule#location AppsecWafCustomRule#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// The status code to return when the WAF custom rule triggers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/appsec_waf_custom_rule#status_code AppsecWafCustomRule#status_code}
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

