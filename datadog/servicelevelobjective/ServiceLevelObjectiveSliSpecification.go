// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecification struct {
	// time_slice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/service_level_objective#time_slice ServiceLevelObjective#time_slice}
	TimeSlice *ServiceLevelObjectiveSliSpecificationTimeSlice `field:"required" json:"timeSlice" yaml:"timeSlice"`
}

