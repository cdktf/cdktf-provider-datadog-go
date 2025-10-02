// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaNumberFormat struct {
	// unit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#unit Dashboard#unit}
	Unit *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaNumberFormatUnit `field:"required" json:"unit" yaml:"unit"`
	// unit_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/dashboard#unit_scale Dashboard#unit_scale}
	UnitScale *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinitionRequestFormulaNumberFormatUnitScale `field:"optional" json:"unitScale" yaml:"unitScale"`
}

