// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryTableDefinitionRequestFormula struct {
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#formula_expression Powerpack#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#alias Powerpack#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A list of display modes for each table cell. Valid values are `number`, `bar`, `trend`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#cell_display_mode Powerpack#cell_display_mode}
	CellDisplayMode *string `field:"optional" json:"cellDisplayMode" yaml:"cellDisplayMode"`
	// cell_display_mode_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#cell_display_mode_options Powerpack#cell_display_mode_options}
	CellDisplayModeOptions *PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions `field:"optional" json:"cellDisplayModeOptions" yaml:"cellDisplayModeOptions"`
	// conditional_formats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#conditional_formats Powerpack#conditional_formats}
	ConditionalFormats interface{} `field:"optional" json:"conditionalFormats" yaml:"conditionalFormats"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#limit Powerpack#limit}
	Limit *PowerpackWidgetQueryTableDefinitionRequestFormulaLimit `field:"optional" json:"limit" yaml:"limit"`
	// number_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#number_format Powerpack#number_format}
	NumberFormat *PowerpackWidgetQueryTableDefinitionRequestFormulaNumberFormat `field:"optional" json:"numberFormat" yaml:"numberFormat"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/powerpack#style Powerpack#style}
	Style *PowerpackWidgetQueryTableDefinitionRequestFormulaStyle `field:"optional" json:"style" yaml:"style"`
}

