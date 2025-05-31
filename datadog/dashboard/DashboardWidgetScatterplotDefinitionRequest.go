// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetScatterplotDefinitionRequest struct {
	// scatterplot_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#scatterplot_table Dashboard#scatterplot_table}
	ScatterplotTable interface{} `field:"optional" json:"scatterplotTable" yaml:"scatterplotTable"`
	// x block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#x Dashboard#x}
	X interface{} `field:"optional" json:"x" yaml:"x"`
	// y block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#y Dashboard#y}
	Y interface{} `field:"optional" json:"y" yaml:"y"`
}

