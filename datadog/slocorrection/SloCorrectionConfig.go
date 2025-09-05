// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package slocorrection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SloCorrectionConfig struct {
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
	// Category the SLO correction belongs to. Valid values are `Scheduled Maintenance`, `Outside Business Hours`, `Deployment`, `Other`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#category SloCorrection#category}
	Category *string `field:"required" json:"category" yaml:"category"`
	// ID of the SLO that this correction will be applied to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#slo_id SloCorrection#slo_id}
	SloId *string `field:"required" json:"sloId" yaml:"sloId"`
	// Starting time of the correction in epoch seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#start SloCorrection#start}
	Start *float64 `field:"required" json:"start" yaml:"start"`
	// Description of the correction being made.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#description SloCorrection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Length of time in seconds for a specified `rrule` recurring SLO correction (required if specifying `rrule`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#duration SloCorrection#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Ending time of the correction in epoch seconds. Required for one time corrections, but optional if `rrule` is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#end SloCorrection#end}
	End *float64 `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#id SloCorrection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Recurrence rules as defined in the iCalendar RFC 5545.
	//
	// Supported rules for SLO corrections are `FREQ`, `INTERVAL`, `COUNT` and `UNTIL`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#rrule SloCorrection#rrule}
	Rrule *string `field:"optional" json:"rrule" yaml:"rrule"`
	// The timezone to display in the UI for the correction times (defaults to "UTC").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/slo_correction#timezone SloCorrection#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

