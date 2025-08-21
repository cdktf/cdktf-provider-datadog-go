// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetTimeseriesDefinitionRequestFormulaNumberFormatUnit struct {
	// canonical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard#canonical Dashboard#canonical}
	Canonical *DashboardWidgetTimeseriesDefinitionRequestFormulaNumberFormatUnitCanonical `field:"optional" json:"canonical" yaml:"canonical"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard#custom Dashboard#custom}
	Custom *DashboardWidgetTimeseriesDefinitionRequestFormulaNumberFormatUnitCustom `field:"optional" json:"custom" yaml:"custom"`
}

