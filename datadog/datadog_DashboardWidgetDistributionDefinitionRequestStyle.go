// Prebuilt datadog Provider for Terraform CDK (cdktf)
package datadog


type DashboardWidgetDistributionDefinitionRequestStyle struct {
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#palette Dashboard#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
}

