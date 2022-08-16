// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetTimeseriesDefinitionMarker struct {
	// A mathematical expression describing the marker, for example: `y > 1`, `-5 < y < 0`, `y = 19`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#value Dashboard#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// How the marker lines are displayed, options are one of {`error`, `warning`, `info`, `ok`} combined with one of {`dashed`, `solid`, `bold`}.
	//
	// Example: `error dashed`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#display_type Dashboard#display_type}
	DisplayType *string `field:"optional" json:"displayType" yaml:"displayType"`
	// A label for the line or range.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#label Dashboard#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

