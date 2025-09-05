// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecificationTimeSliceQueryFormula struct {
	// The formula string, which is an expression involving named queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/service_level_objective#formula_expression ServiceLevelObjective#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
}

