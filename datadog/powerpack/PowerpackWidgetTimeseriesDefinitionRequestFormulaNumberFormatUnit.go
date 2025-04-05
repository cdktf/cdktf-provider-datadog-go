// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionRequestFormulaNumberFormatUnit struct {
	// canonical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/powerpack#canonical Powerpack#canonical}
	Canonical *PowerpackWidgetTimeseriesDefinitionRequestFormulaNumberFormatUnitCanonical `field:"optional" json:"canonical" yaml:"canonical"`
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/powerpack#custom Powerpack#custom}
	Custom *PowerpackWidgetTimeseriesDefinitionRequestFormulaNumberFormatUnitCustom `field:"optional" json:"custom" yaml:"custom"`
}

