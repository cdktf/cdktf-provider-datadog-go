package dashboard


type DashboardWidgetIframeDefinition struct {
	// The URL to use as a data source for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.26.0/docs/resources/dashboard#url Dashboard#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

