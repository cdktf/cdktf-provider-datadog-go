// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetGeomapDefinitionStyle struct {
	// The color palette to apply to the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#palette Powerpack#palette}
	Palette *string `field:"required" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#palette_flip Powerpack#palette_flip}
	PaletteFlip interface{} `field:"required" json:"paletteFlip" yaml:"paletteFlip"`
}

