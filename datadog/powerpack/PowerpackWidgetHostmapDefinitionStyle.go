// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetHostmapDefinitionStyle struct {
	// The max value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#fill_max Powerpack#fill_max}
	FillMax *string `field:"optional" json:"fillMax" yaml:"fillMax"`
	// The min value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#fill_min Powerpack#fill_min}
	FillMin *string `field:"optional" json:"fillMin" yaml:"fillMin"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#palette Powerpack#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#palette_flip Powerpack#palette_flip}
	PaletteFlip interface{} `field:"optional" json:"paletteFlip" yaml:"paletteFlip"`
}

