// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafcustomrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppsecWafCustomRuleConfig struct {
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
	// Indicates whether the WAF custom rule will block the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#blocking AppsecWafCustomRule#blocking}
	Blocking interface{} `field:"required" json:"blocking" yaml:"blocking"`
	// Indicates whether the WAF custom rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#enabled AppsecWafCustomRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The Name of the WAF custom rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#name AppsecWafCustomRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags associated with the WAF custom rule.
	//
	// `category` and `type` tags are required. Supported categories include `business_logic`, `attack_attempt` and `security_response`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#tags AppsecWafCustomRule#tags}
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#action AppsecWafCustomRule#action}
	Action *AppsecWafCustomRuleAction `field:"optional" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#condition AppsecWafCustomRule#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// The path glob for the WAF custom rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#path_glob AppsecWafCustomRule#path_glob}
	PathGlob *string `field:"optional" json:"pathGlob" yaml:"pathGlob"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/appsec_waf_custom_rule#scope AppsecWafCustomRule#scope}
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
}

