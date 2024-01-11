// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetScatterplotDefinitionRequest struct {
	// scatterplot_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/powerpack#scatterplot_table Powerpack#scatterplot_table}
	ScatterplotTable interface{} `field:"optional" json:"scatterplotTable" yaml:"scatterplotTable"`
	// x block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/powerpack#x Powerpack#x}
	X interface{} `field:"optional" json:"x" yaml:"x"`
	// y block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/resources/powerpack#y Powerpack#y}
	Y interface{} `field:"optional" json:"y" yaml:"y"`
}

