// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringdefaultrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityMonitoringDefaultRuleConfig struct {
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
	// case block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#case SecurityMonitoringDefaultRule#case}
	Case interface{} `field:"optional" json:"case" yaml:"case"`
	// Custom Message (will override default message) for generated signals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#custom_message SecurityMonitoringDefaultRule#custom_message}
	CustomMessage *string `field:"optional" json:"customMessage" yaml:"customMessage"`
	// The name (will override default name) of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#custom_name SecurityMonitoringDefaultRule#custom_name}
	CustomName *string `field:"optional" json:"customName" yaml:"customName"`
	// Custom tags for generated signals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#custom_tags SecurityMonitoringDefaultRule#custom_tags}
	CustomTags *[]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Enable the rule. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#enabled SecurityMonitoringDefaultRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#filter SecurityMonitoringDefaultRule#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#id SecurityMonitoringDefaultRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#options SecurityMonitoringDefaultRule#options}
	Options *SecurityMonitoringDefaultRuleOptions `field:"optional" json:"options" yaml:"options"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/security_monitoring_default_rule#query SecurityMonitoringDefaultRule#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
}

