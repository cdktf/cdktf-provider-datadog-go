// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetToplistDefinitionRequestFormulaNumberFormatUnitCanonical struct {
	// Unit name. It should be in singular form ('megabyte' and not 'megabytes').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/powerpack#unit_name Powerpack#unit_name}
	UnitName *string `field:"required" json:"unitName" yaml:"unitName"`
	// per unit name. If you want to represent megabytes/s, you set 'unit_name' = 'megabyte' and 'per_unit_name = 'second'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/powerpack#per_unit_name Powerpack#per_unit_name}
	PerUnitName *string `field:"optional" json:"perUnitName" yaml:"perUnitName"`
}

