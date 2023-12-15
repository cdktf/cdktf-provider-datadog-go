// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetHostmapDefinitionStyle struct {
	// The max value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/dashboard#fill_max Dashboard#fill_max}
	FillMax *string `field:"optional" json:"fillMax" yaml:"fillMax"`
	// The min value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/dashboard#fill_min Dashboard#fill_min}
	FillMin *string `field:"optional" json:"fillMin" yaml:"fillMin"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/dashboard#palette_flip Dashboard#palette_flip}
	PaletteFlip interface{} `field:"optional" json:"paletteFlip" yaml:"paletteFlip"`
}

