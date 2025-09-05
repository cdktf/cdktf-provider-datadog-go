// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnit struct {
	// canonical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#canonical Dashboard#canonical}
	Canonical *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitCanonical `field:"optional" json:"canonical" yaml:"canonical"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#custom Dashboard#custom}
	Custom *DashboardWidgetTreemapDefinitionRequestFormulaNumberFormatUnitCustom `field:"optional" json:"custom" yaml:"custom"`
}

