// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetFreeTextDefinition struct {
	// The text to display in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/powerpack#text Powerpack#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// The color of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/powerpack#color Powerpack#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The size of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/powerpack#font_size Powerpack#font_size}
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// The alignment of the text in the widget. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/powerpack#text_align Powerpack#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
}

