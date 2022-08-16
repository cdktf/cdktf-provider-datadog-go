// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetScatterplotDefinitionRequestScatterplotTableFormula struct {
	// Dimension of the Scatterplot. Valid values are `x`, `y`, `radius`, `color`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#dimension Dashboard#dimension}
	Dimension *string `field:"required" json:"dimension" yaml:"dimension"`
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#formula_expression Dashboard#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#alias Dashboard#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

