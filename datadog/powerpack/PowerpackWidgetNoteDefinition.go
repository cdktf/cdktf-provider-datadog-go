// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetNoteDefinition struct {
	// The content of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#content Powerpack#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The background color of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#background_color Powerpack#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The size of the text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#font_size Powerpack#font_size}
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Whether to add padding or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#has_padding Powerpack#has_padding}
	HasPadding interface{} `field:"optional" json:"hasPadding" yaml:"hasPadding"`
	// Whether to show a tick or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#show_tick Powerpack#show_tick}
	ShowTick interface{} `field:"optional" json:"showTick" yaml:"showTick"`
	// The alignment of the widget's text. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#text_align Powerpack#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
	// When `tick = true`, a string indicating on which side of the widget the tick should be displayed.
	//
	// Valid values are `bottom`, `left`, `right`, `top`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#tick_edge Powerpack#tick_edge}
	TickEdge *string `field:"optional" json:"tickEdge" yaml:"tickEdge"`
	// When `tick = true`, a string with a percent sign indicating the position of the tick, for example: `tick_pos = "50%"` is centered alignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#tick_pos Powerpack#tick_pos}
	TickPos *string `field:"optional" json:"tickPos" yaml:"tickPos"`
	// The vertical alignment for the widget. Valid values are `center`, `top`, `bottom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/powerpack#vertical_align Powerpack#vertical_align}
	VerticalAlign *string `field:"optional" json:"verticalAlign" yaml:"verticalAlign"`
}

