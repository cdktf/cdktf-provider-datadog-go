// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetListStreamDefinitionRequestColumns struct {
	// Widget column field.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#field Dashboard#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Widget column width. Valid values are `auto`, `compact`, `full`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#width Dashboard#width}
	Width *string `field:"required" json:"width" yaml:"width"`
}

