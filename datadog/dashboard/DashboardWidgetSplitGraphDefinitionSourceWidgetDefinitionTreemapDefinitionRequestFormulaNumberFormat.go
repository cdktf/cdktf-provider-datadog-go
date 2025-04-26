// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestFormulaNumberFormat struct {
	// unit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/dashboard#unit Dashboard#unit}
	Unit *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestFormulaNumberFormatUnit `field:"required" json:"unit" yaml:"unit"`
	// unit_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/dashboard#unit_scale Dashboard#unit_scale}
	UnitScale *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestFormulaNumberFormatUnitScale `field:"optional" json:"unitScale" yaml:"unitScale"`
}

