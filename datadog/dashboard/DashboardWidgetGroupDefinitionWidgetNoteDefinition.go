// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetNoteDefinition struct {
	// The content of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#content Dashboard#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The background color of the note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#background_color Dashboard#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// The size of the text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#font_size Dashboard#font_size}
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// Whether to add padding or not. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#has_padding Dashboard#has_padding}
	HasPadding interface{} `field:"optional" json:"hasPadding" yaml:"hasPadding"`
	// Whether to show a tick or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#show_tick Dashboard#show_tick}
	ShowTick interface{} `field:"optional" json:"showTick" yaml:"showTick"`
	// The alignment of the widget's text. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#text_align Dashboard#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
	// When `tick = true`, a string indicating on which side of the widget the tick should be displayed.
	//
	// Valid values are `bottom`, `left`, `right`, `top`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#tick_edge Dashboard#tick_edge}
	TickEdge *string `field:"optional" json:"tickEdge" yaml:"tickEdge"`
	// When `tick = true`, a string with a percent sign indicating the position of the tick, for example: `tick_pos = "50%"` is centered alignment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#tick_pos Dashboard#tick_pos}
	TickPos *string `field:"optional" json:"tickPos" yaml:"tickPos"`
	// The vertical alignment for the widget. Valid values are `center`, `top`, `bottom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#vertical_align Dashboard#vertical_align}
	VerticalAlign *string `field:"optional" json:"verticalAlign" yaml:"verticalAlign"`
}

