// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetGeomapDefinitionRequestFormula struct {
	// A string expression built from queries, formulas, and functions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#formula_expression Dashboard#formula_expression}
	FormulaExpression *string `field:"required" json:"formulaExpression" yaml:"formulaExpression"`
	// An expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#alias Dashboard#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A list of display modes for each table cell. Valid values are `number`, `bar`, `trend`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#cell_display_mode Dashboard#cell_display_mode}
	CellDisplayMode *string `field:"optional" json:"cellDisplayMode" yaml:"cellDisplayMode"`
	// conditional_formats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#conditional_formats Dashboard#conditional_formats}
	ConditionalFormats interface{} `field:"optional" json:"conditionalFormats" yaml:"conditionalFormats"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#limit Dashboard#limit}
	Limit *DashboardWidgetGroupDefinitionWidgetGeomapDefinitionRequestFormulaLimit `field:"optional" json:"limit" yaml:"limit"`
	// number_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#number_format Dashboard#number_format}
	NumberFormat *DashboardWidgetGroupDefinitionWidgetGeomapDefinitionRequestFormulaNumberFormat `field:"optional" json:"numberFormat" yaml:"numberFormat"`
	// style block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/dashboard#style Dashboard#style}
	Style *DashboardWidgetGroupDefinitionWidgetGeomapDefinitionRequestFormulaStyle `field:"optional" json:"style" yaml:"style"`
}

