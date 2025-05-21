// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitynotificationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityNotificationRuleConfig struct {
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
	// The name of the rule (must be unique).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/security_notification_rule#name SecurityNotificationRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of handle targets for the notifications.
	//
	// A target must be prefixed with an.
	Targets *[]*string `field:"required" json:"targets" yaml:"targets"`
	// Indicates whether the rule is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/security_notification_rule#enabled SecurityNotificationRule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/security_notification_rule#selectors SecurityNotificationRule#selectors}
	Selectors *SecurityNotificationRuleSelectors `field:"optional" json:"selectors" yaml:"selectors"`
	// Specifies the time period, in seconds, used to aggregate the notification. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/security_notification_rule#time_aggregation SecurityNotificationRule#time_aggregation}
	TimeAggregation *float64 `field:"optional" json:"timeAggregation" yaml:"timeAggregation"`
}

