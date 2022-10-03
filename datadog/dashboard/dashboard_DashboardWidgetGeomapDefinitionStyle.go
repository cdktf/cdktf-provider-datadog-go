package dashboard


type DashboardWidgetGeomapDefinitionStyle struct {
	// The color palette to apply to the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#palette Dashboard#palette}
	Palette *string `field:"required" json:"palette" yaml:"palette"`
	// A Boolean indicating whether to flip the palette tones.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#palette_flip Dashboard#palette_flip}
	PaletteFlip interface{} `field:"required" json:"paletteFlip" yaml:"paletteFlip"`
}

