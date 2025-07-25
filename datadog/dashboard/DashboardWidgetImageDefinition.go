// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetImageDefinition struct {
	// The URL to use as a data source for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#url Dashboard#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Whether to display a background or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#has_background Dashboard#has_background}
	HasBackground interface{} `field:"optional" json:"hasBackground" yaml:"hasBackground"`
	// Whether to display a border or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#has_border Dashboard#has_border}
	HasBorder interface{} `field:"optional" json:"hasBorder" yaml:"hasBorder"`
	// The horizontal alignment for the widget. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#horizontal_align Dashboard#horizontal_align}
	HorizontalAlign *string `field:"optional" json:"horizontalAlign" yaml:"horizontalAlign"`
	// The margins to use around the image.
	//
	// Note: `small` and `large` values are deprecated. Valid values are `sm`, `md`, `lg`, `small`, `large`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#margin Dashboard#margin}
	Margin *string `field:"optional" json:"margin" yaml:"margin"`
	// The preferred method to adapt the dimensions of the image.
	//
	// The values are based on the image `object-fit` CSS properties. Note: `zoom`, `fit` and `center` values are deprecated. Valid values are `fill`, `contain`, `cover`, `none`, `scale-down`, `zoom`, `fit`, `center`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#sizing Dashboard#sizing}
	Sizing *string `field:"optional" json:"sizing" yaml:"sizing"`
	// The URL in dark mode to use as a data source for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#url_dark_theme Dashboard#url_dark_theme}
	UrlDarkTheme *string `field:"optional" json:"urlDarkTheme" yaml:"urlDarkTheme"`
	// The vertical alignment for the widget. Valid values are `center`, `top`, `bottom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#vertical_align Dashboard#vertical_align}
	VerticalAlign *string `field:"optional" json:"verticalAlign" yaml:"verticalAlign"`
}

