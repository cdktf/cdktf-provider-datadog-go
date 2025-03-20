// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafexclusionfilter


type AppsecWafExclusionFilterRulesTarget struct {
	// Target a single WAF rule based on its identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#rule_id AppsecWafExclusionFilter#rule_id}
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#tags AppsecWafExclusionFilter#tags}
	Tags *AppsecWafExclusionFilterRulesTargetTags `field:"optional" json:"tags" yaml:"tags"`
}

