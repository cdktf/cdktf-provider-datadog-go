// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecificationTimeSliceQuery struct {
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/service_level_objective#formula ServiceLevelObjective#formula}
	Formula *ServiceLevelObjectiveSliSpecificationTimeSliceQueryFormula `field:"required" json:"formula" yaml:"formula"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/service_level_objective#query ServiceLevelObjective#query}
	Query interface{} `field:"required" json:"query" yaml:"query"`
}

