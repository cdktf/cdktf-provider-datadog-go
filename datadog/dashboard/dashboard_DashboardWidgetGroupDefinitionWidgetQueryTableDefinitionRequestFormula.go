package dashboard


type DashboardWidgetGroupDefinitionWidgetQueryTableDefinitionRequestFormula struct {
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#formula_expression Dashboard#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#alias Dashboard#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A list of display modes for each table cell. Valid values are `number`, `bar`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#cell_display_mode Dashboard#cell_display_mode}
	CellDisplayMode *string `field:"optional" json:"cellDisplayMode" yaml:"cellDisplayMode"`
	// conditional_formats block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#conditional_formats Dashboard#conditional_formats}
	ConditionalFormats interface{} `field:"optional" json:"conditionalFormats" yaml:"conditionalFormats"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#limit Dashboard#limit}
	Limit *DashboardWidgetGroupDefinitionWidgetQueryTableDefinitionRequestFormulaLimit `field:"optional" json:"limit" yaml:"limit"`
}

