// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsecwafexclusionfilter


type AppsecWafExclusionFilterScope struct {
	// Deploy on this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/appsec_waf_exclusion_filter#env AppsecWafExclusionFilter#env}
	Env *string `field:"optional" json:"env" yaml:"env"`
	// Deploy on this service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/appsec_waf_exclusion_filter#service AppsecWafExclusionFilter#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

