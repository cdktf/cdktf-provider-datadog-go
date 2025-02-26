// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetTimeseriesDefinitionRequestFormulaNumberFormatUnitCanonical struct {
	// Unit name. It should be in singular form ('megabyte' and not 'megabytes').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#unit_name Dashboard#unit_name}
	UnitName *string `field:"required" json:"unitName" yaml:"unitName"`
	// per unit name. If you want to represent megabytes/s, you set 'unit_name' = 'megabyte' and 'per_unit_name = 'second'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#per_unit_name Dashboard#per_unit_name}
	PerUnitName *string `field:"optional" json:"perUnitName" yaml:"perUnitName"`
}

