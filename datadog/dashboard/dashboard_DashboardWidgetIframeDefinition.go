package dashboard


type DashboardWidgetIframeDefinition struct {
	// The URL to use as a data source for the widget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#url Dashboard#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}
