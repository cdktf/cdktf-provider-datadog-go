// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetScatterplotDefinitionRequestScatterplotTableFormula struct {
	// Dimension of the Scatterplot. Valid values are `x`, `y`, `radius`, `color`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#dimension Powerpack#dimension}
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#formula_expression Powerpack#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#alias Powerpack#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

