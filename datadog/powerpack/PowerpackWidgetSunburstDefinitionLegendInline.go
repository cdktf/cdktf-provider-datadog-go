// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetSunburstDefinitionLegendInline struct {
	// The type of legend (inline or automatic). Valid values are `inline`, `automatic`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/powerpack#type Powerpack#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether to hide the percentages of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/powerpack#hide_percent Powerpack#hide_percent}
	HidePercent interface{} `field:"optional" json:"hidePercent" yaml:"hidePercent"`
	// Whether to hide the values of the groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/powerpack#hide_value Powerpack#hide_value}
	HideValue interface{} `field:"optional" json:"hideValue" yaml:"hideValue"`
}

