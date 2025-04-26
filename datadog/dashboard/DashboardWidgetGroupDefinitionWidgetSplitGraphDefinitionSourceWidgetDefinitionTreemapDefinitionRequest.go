// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequest struct {
	// formula block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/dashboard#formula Dashboard#formula}
	Formula interface{} `field:"optional" json:"formula" yaml:"formula"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/dashboard#query Dashboard#query}
	Query interface{} `field:"optional" json:"query" yaml:"query"`
}

