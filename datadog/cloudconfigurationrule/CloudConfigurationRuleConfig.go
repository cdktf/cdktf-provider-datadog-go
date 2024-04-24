// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudconfigurationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudConfigurationRuleConfig struct {
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
	// Whether the cloud configuration rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#enabled CloudConfigurationRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The message associated to the rule that will be shown in findings and signals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#message CloudConfigurationRule#message}
	Message *string `field:"required" json:"message" yaml:"message"`
	// The name of the cloud configuration rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#name CloudConfigurationRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policy written in Rego format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#policy CloudConfigurationRule#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Main resource type to be checked by the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#resource_type CloudConfigurationRule#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Severity of the rule and associated signals. Valid values are `info`, `low`, `medium`, `high`, `critical`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#severity CloudConfigurationRule#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#filter CloudConfigurationRule#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Defaults to empty list.
	//
	// This function will be deprecated soon. Use the notification rules function instead. Fields to group by when generating signals, e.g.
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#id CloudConfigurationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This function will be deprecated soon.
	//
	// Use the notification rules function instead. Notification targets for signals. Defaults to empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#notifications CloudConfigurationRule#notifications}
	Notifications *[]*string `field:"optional" json:"notifications" yaml:"notifications"`
	// Related resource types to be checked by the rule. Defaults to empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#related_resource_types CloudConfigurationRule#related_resource_types}
	RelatedResourceTypes *[]*string `field:"optional" json:"relatedResourceTypes" yaml:"relatedResourceTypes"`
	// Tags of the rule, propagated to findings and signals. Defaults to empty list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/cloud_configuration_rule#tags CloudConfigurationRule#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

