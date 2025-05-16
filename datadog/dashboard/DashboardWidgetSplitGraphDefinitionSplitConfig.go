// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSplitConfig struct {
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/dashboard#sort Dashboard#sort}
	Sort *DashboardWidgetSplitGraphDefinitionSplitConfigSort `field:"required" json:"sort" yaml:"sort"`
	// split_dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/dashboard#split_dimensions Dashboard#split_dimensions}
	SplitDimensions *DashboardWidgetSplitGraphDefinitionSplitConfigSplitDimensions `field:"required" json:"splitDimensions" yaml:"splitDimensions"`
	// Maximum number of graphs to display in the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// static_splits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.62.0/docs/resources/dashboard#static_splits Dashboard#static_splits}
	StaticSplits interface{} `field:"optional" json:"staticSplits" yaml:"staticSplits"`
}

