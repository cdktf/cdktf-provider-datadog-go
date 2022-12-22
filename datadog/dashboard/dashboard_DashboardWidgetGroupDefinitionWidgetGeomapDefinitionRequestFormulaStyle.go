package dashboard


type DashboardWidgetGroupDefinitionWidgetGeomapDefinitionRequestFormulaStyle struct {
	// The color palette used to display the formula.
	//
	// A guide to the available color palettes can be found at https://docs.datadoghq.com/dashboards/guide/widget_colors
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// Index specifying which color to use within the palette.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#palette_index Dashboard#palette_index}
	PaletteIndex *float64 `field:"optional" json:"paletteIndex" yaml:"paletteIndex"`
}

