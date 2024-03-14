// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitymonitoringrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityMonitoringRuleConfig struct {
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
	// Message for generated signals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#message SecurityMonitoringRule#message}
	Message *string `field:"required" json:"message" yaml:"message"`
	// The name of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#name SecurityMonitoringRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// case block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#case SecurityMonitoringRule#case}
	Case interface{} `field:"optional" json:"case" yaml:"case"`
	// Whether the rule is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#enabled SecurityMonitoringRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#filter SecurityMonitoringRule#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Whether the notifications include the triggering group-by values in their title. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#has_extended_title SecurityMonitoringRule#has_extended_title}
	HasExtendedTitle interface{} `field:"optional" json:"hasExtendedTitle" yaml:"hasExtendedTitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#id SecurityMonitoringRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#options SecurityMonitoringRule#options}
	Options *SecurityMonitoringRuleOptions `field:"optional" json:"options" yaml:"options"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#query SecurityMonitoringRule#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
	// signal_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#signal_query SecurityMonitoringRule#signal_query}
	SignalQuery interface{} `field:"optional" json:"signalQuery" yaml:"signalQuery"`
	// Tags for generated signals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#tags SecurityMonitoringRule#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// third_party_case block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#third_party_case SecurityMonitoringRule#third_party_case}
	ThirdPartyCase interface{} `field:"optional" json:"thirdPartyCase" yaml:"thirdPartyCase"`
	// The rule type. Valid values are `application_security`, `log_detection`, `workload_security`, `signal_correlation`. Defaults to `"log_detection"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/security_monitoring_rule#type SecurityMonitoringRule#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

