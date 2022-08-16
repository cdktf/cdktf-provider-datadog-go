// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetGeomapDefinitionView struct {
	// The two-letter ISO code of a country to focus the map on (or `WORLD`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#focus Dashboard#focus}
	Focus *string `field:"required" json:"focus" yaml:"focus"`
}

