// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetTimeseriesDefinitionYaxis struct {
	// Always include zero or fit the axis to the data range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#include_zero Dashboard#include_zero}
	IncludeZero interface{} `field:"optional" json:"includeZero" yaml:"includeZero"`
	// The label of the axis to display on the graph.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#label Dashboard#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Specify the maximum value to show on the Y-axis.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#max Dashboard#max}
	Max *string `field:"optional" json:"max" yaml:"max"`
	// Specify the minimum value to show on the Y-axis.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#min Dashboard#min}
	Min *string `field:"optional" json:"min" yaml:"min"`
	// Specify the scale type, options: `linear`, `log`, `pow`, `sqrt`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#scale Dashboard#scale}
	Scale *string `field:"optional" json:"scale" yaml:"scale"`
}
