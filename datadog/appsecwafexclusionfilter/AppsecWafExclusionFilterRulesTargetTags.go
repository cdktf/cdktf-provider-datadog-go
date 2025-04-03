// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafexclusionfilter


type AppsecWafExclusionFilterRulesTargetTags struct {
	// The category of the targeted WAF rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_exclusion_filter#category AppsecWafExclusionFilter#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The type of the targeted WAF rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/appsec_waf_exclusion_filter#type AppsecWafExclusionFilter#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

