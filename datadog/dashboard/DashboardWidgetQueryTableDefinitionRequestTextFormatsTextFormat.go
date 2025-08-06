// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetQueryTableDefinitionRequestTextFormatsTextFormat struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/dashboard#match Dashboard#match}
	Match *DashboardWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch `field:"required" json:"match" yaml:"match"`
	// The custom color palette to apply to the background.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/dashboard#custom_bg_color Dashboard#custom_bg_color}
	CustomBgColor *string `field:"optional" json:"customBgColor" yaml:"customBgColor"`
	// The custom color palette to apply to the foreground text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/dashboard#custom_fg_color Dashboard#custom_fg_color}
	CustomFgColor *string `field:"optional" json:"customFgColor" yaml:"customFgColor"`
	// The color palette to apply. Valid values are `white_on_red`, `white_on_yellow`, `white_on_green`, `black_on_light_red`, `black_on_light_yellow`, `black_on_light_green`, `red_on_white`, `yellow_on_white`, `green_on_white`, `custom_bg`, `custom_text`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// replace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/dashboard#replace Dashboard#replace}
	Replace *DashboardWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace `field:"optional" json:"replace" yaml:"replace"`
}

