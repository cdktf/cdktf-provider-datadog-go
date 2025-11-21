// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oncallschedule


type OnCallScheduleLayer struct {
	// The date/time when this layer should become active (in ISO 8601).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_schedule#effective_date OnCallSchedule#effective_date}
	EffectiveDate *string `field:"required" json:"effectiveDate" yaml:"effectiveDate"`
	// The name of this layer. Should be unique within the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_schedule#name OnCallSchedule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The date/time when the rotation for this layer starts (in ISO 8601).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_schedule#rotation_start OnCallSchedule#rotation_start}
	RotationStart *string `field:"required" json:"rotationStart" yaml:"rotationStart"`
	// List of user IDs for the layer. Can either be a valid user id or null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_schedule#users OnCallSchedule#users}
	Users *[]*string `field:"required" json:"users" yaml:"users"`
	// The date/time after which this layer no longer applies (in ISO 8601).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_schedule#end_date OnCallSchedule#end_date}
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// interval block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_schedule#interval OnCallSchedule#interval}
	Interval *OnCallScheduleLayerInterval `field:"optional" json:"interval" yaml:"interval"`
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/on_call_schedule#restriction OnCallSchedule#restriction}
	Restriction interface{} `field:"optional" json:"restriction" yaml:"restriction"`
}

