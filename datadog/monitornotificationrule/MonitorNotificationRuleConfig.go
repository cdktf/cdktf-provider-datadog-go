// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitornotificationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorNotificationRuleConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/monitor_notification_rule#filter MonitorNotificationRule#filter}
	Filter *MonitorNotificationRuleFilter `field:"required" json:"filter" yaml:"filter"`
	// The name of the monitor notification rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/monitor_notification_rule#name MonitorNotificationRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of recipients to notify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/monitor_notification_rule#recipients MonitorNotificationRule#recipients}
	Recipients *[]*string `field:"required" json:"recipients" yaml:"recipients"`
}

