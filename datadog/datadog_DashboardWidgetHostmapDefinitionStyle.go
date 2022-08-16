// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetHostmapDefinitionStyle struct {
	// The max value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#fill_max Dashboard#fill_max}
	FillMax *string `field:"optional" json:"fillMax" yaml:"fillMax"`
	// The min value to use to color the map.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#fill_min Dashboard#fill_min}
	FillMin *string `field:"optional" json:"fillMin" yaml:"fillMin"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#palette_flip Dashboard#palette_flip}
	PaletteFlip interface{} `field:"optional" json:"paletteFlip" yaml:"paletteFlip"`
}

