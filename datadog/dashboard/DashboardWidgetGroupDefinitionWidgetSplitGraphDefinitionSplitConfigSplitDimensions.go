// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSplitConfigSplitDimensions struct {
	// The system interprets this attribute differently depending on the data source of the query being split.
	//
	// For metrics, it's a tag. For the events platform, it's an attribute or tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.1/docs/resources/dashboard#one_graph_per Dashboard#one_graph_per}
	OneGraphPer *string `field:"required" json:"oneGraphPer" yaml:"oneGraphPer"`
}

