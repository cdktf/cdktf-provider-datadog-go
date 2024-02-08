// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorconfigpolicy


type MonitorConfigPolicyTagPolicy struct {
	// The key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.0/docs/resources/monitor_config_policy#tag_key MonitorConfigPolicy#tag_key}
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// If a tag key is required for monitor creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.0/docs/resources/monitor_config_policy#tag_key_required MonitorConfigPolicy#tag_key_required}
	TagKeyRequired interface{} `field:"required" json:"tagKeyRequired" yaml:"tagKeyRequired"`
	// Valid values for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.0/docs/resources/monitor_config_policy#valid_tag_values MonitorConfigPolicy#valid_tag_values}
	ValidTagValues *[]*string `field:"required" json:"validTagValues" yaml:"validTagValues"`
}

