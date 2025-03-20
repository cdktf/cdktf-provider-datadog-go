// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafexclusionfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsecWafExclusionFilterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A description for the exclusion filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#description AppsecWafExclusionFilter#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Indicates whether the exclusion filter is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#enabled AppsecWafExclusionFilter#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The event query matched by the legacy exclusion filter. Cannot be created nor updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#event_query AppsecWafExclusionFilter#event_query}
	EventQuery *string `field:"optional" json:"eventQuery" yaml:"eventQuery"`
	// The client IP addresses matched by the exclusion filter (CIDR notation is supported).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#ip_list AppsecWafExclusionFilter#ip_list}
	IpList *[]*string `field:"optional" json:"ipList" yaml:"ipList"`
	// The action taken when the exclusion filter matches.
	//
	// When set to `monitor`, security traces are emitted but the requests are not blocked. By default, security traces are not emitted and the requests are not blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#on_match AppsecWafExclusionFilter#on_match}
	OnMatch *string `field:"optional" json:"onMatch" yaml:"onMatch"`
	// A list of parameters matched by the exclusion filter in the HTTP query string and HTTP request body.
	//
	// Nested parameters can be matched by joining fields with a dot character.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#parameters AppsecWafExclusionFilter#parameters}
	Parameters *[]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The HTTP path glob expression matched by the exclusion filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#path_glob AppsecWafExclusionFilter#path_glob}
	PathGlob *string `field:"optional" json:"pathGlob" yaml:"pathGlob"`
	// rules_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#rules_target AppsecWafExclusionFilter#rules_target}
	RulesTarget interface{} `field:"optional" json:"rulesTarget" yaml:"rulesTarget"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/appsec_waf_exclusion_filter#scope AppsecWafExclusionFilter#scope}
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
}

