// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnit struct {
	// canonical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/dashboard#canonical Dashboard#canonical}
	Canonical *DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCanonical `field:"optional" json:"canonical" yaml:"canonical"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/dashboard#custom Dashboard#custom}
	Custom *DashboardWidgetGeomapDefinitionRequestFormulaNumberFormatUnitCustom `field:"optional" json:"custom" yaml:"custom"`
}

