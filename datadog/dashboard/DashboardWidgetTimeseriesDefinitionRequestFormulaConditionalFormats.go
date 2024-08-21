// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetTimeseriesDefinitionRequestFormulaConditionalFormats struct {
	// The comparator to use. Valid values are `=`, `>`, `>=`, `<`, `<=`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#comparator Dashboard#comparator}
	Comparator *string `field:"required" json:"comparator" yaml:"comparator"`
	// The color palette to apply.
	//
	// Valid values are `blue`, `custom_bg`, `custom_image`, `custom_text`, `gray_on_white`, `grey`, `green`, `orange`, `red`, `red_on_white`, `white_on_gray`, `white_on_green`, `green_on_white`, `white_on_red`, `white_on_yellow`, `yellow_on_white`, `black_on_light_yellow`, `black_on_light_green`, `black_on_light_red`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#palette Dashboard#palette}
	Palette *string `field:"required" json:"palette" yaml:"palette"`
	// A value for the comparator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#value Dashboard#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// The color palette to apply to the background, same values available as palette.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#custom_bg_color Dashboard#custom_bg_color}
	CustomBgColor *string `field:"optional" json:"customBgColor" yaml:"customBgColor"`
	// The color palette to apply to the foreground, same values available as palette.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#custom_fg_color Dashboard#custom_fg_color}
	CustomFgColor *string `field:"optional" json:"customFgColor" yaml:"customFgColor"`
	// Setting this to True hides values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#hide_value Dashboard#hide_value}
	HideValue interface{} `field:"optional" json:"hideValue" yaml:"hideValue"`
	// Displays an image as the background.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#image_url Dashboard#image_url}
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// The metric from the request to correlate with this conditional format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#metric Dashboard#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Defines the displayed timeframe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#timeframe Dashboard#timeframe}
	Timeframe *string `field:"optional" json:"timeframe" yaml:"timeframe"`
}

