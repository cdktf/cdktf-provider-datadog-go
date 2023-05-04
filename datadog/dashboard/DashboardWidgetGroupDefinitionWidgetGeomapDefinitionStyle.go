package dashboard


type DashboardWidgetGroupDefinitionWidgetGeomapDefinitionStyle struct {
	// The color palette to apply to the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#palette Dashboard#palette}
	Palette *string `field:"required" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.1/docs/resources/dashboard#palette_flip Dashboard#palette_flip}
	PaletteFlip interface{} `field:"required" json:"paletteFlip" yaml:"paletteFlip"`
}

