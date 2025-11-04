// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryTableDefinitionRequestFormulaNumberFormat struct {
	// unit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#unit Powerpack#unit}
	Unit *PowerpackWidgetQueryTableDefinitionRequestFormulaNumberFormatUnit `field:"required" json:"unit" yaml:"unit"`
	// unit_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#unit_scale Powerpack#unit_scale}
	UnitScale *PowerpackWidgetQueryTableDefinitionRequestFormulaNumberFormatUnitScale `field:"optional" json:"unitScale" yaml:"unitScale"`
}

