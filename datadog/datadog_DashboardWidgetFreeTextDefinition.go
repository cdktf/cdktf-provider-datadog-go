// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetFreeTextDefinition struct {
	// The text to display in the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#text Dashboard#text}
	Text *string `field:"required" json:"text" yaml:"text"`
	// The color of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#color Dashboard#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// The size of the text in the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#font_size Dashboard#font_size}
	FontSize *string `field:"optional" json:"fontSize" yaml:"fontSize"`
	// The alignment of the text in the widget. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#text_align Dashboard#text_align}
	TextAlign *string `field:"optional" json:"textAlign" yaml:"textAlign"`
}
