package dashboard


type DashboardWidgetGroupDefinitionWidgetScatterplotDefinitionRequest struct {
	// scatterplot_table block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#scatterplot_table Dashboard#scatterplot_table}
	ScatterplotTable interface{} `field:"optional" json:"scatterplotTable" yaml:"scatterplotTable"`
	// x block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#x Dashboard#x}
	X interface{} `field:"optional" json:"x" yaml:"x"`
	// y block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/datadog/r/dashboard#y Dashboard#y}
	Y interface{} `field:"optional" json:"y" yaml:"y"`
}

