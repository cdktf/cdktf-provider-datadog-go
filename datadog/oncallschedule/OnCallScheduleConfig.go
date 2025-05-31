// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OnCallScheduleConfig struct {
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
	// A human-readable name for the new schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_schedule#name OnCallSchedule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The time zone in which the schedule is defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_schedule#time_zone OnCallSchedule#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// layer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_schedule#layer OnCallSchedule#layer}
	Layer interface{} `field:"optional" json:"layer" yaml:"layer"`
	// A list of team ids associated with the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/on_call_schedule#teams OnCallSchedule#teams}
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
}

