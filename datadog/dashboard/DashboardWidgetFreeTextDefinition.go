// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetFreeTextDefinition struct {
	// The text to display in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#text Dashboard#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// The color of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#color Dashboard#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The size of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#font_size Dashboard#font_size}
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// The alignment of the text in the widget. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.30.0/docs/resources/dashboard#text_align Dashboard#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
}

