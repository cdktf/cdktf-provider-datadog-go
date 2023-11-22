// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveQuery struct {
	// The sum of the `total` events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/resources/service_level_objective#denominator ServiceLevelObjective#denominator}
	Denominator *string `field:"required" json:"denominator" yaml:"denominator"`
	// The sum of all the `good` events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/resources/service_level_objective#numerator ServiceLevelObjective#numerator}
	Numerator *string `field:"required" json:"numerator" yaml:"numerator"`
}

