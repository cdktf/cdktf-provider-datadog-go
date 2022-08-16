// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetSunburstDefinitionLegendInline struct {
	// The type of legend (inline or automatic). Valid values are `inline`, `automatic`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Whether to hide the percentages of the groups.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#hide_percent Dashboard#hide_percent}
	HidePercent interface{} `field:"optional" json:"hidePercent" yaml:"hidePercent"`
	// Whether to hide the values of the groups.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#hide_value Dashboard#hide_value}
	HideValue interface{} `field:"optional" json:"hideValue" yaml:"hideValue"`
}

