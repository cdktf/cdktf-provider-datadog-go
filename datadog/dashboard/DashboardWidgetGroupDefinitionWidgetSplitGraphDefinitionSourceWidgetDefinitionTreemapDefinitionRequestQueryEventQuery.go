// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/dashboard#compute Dashboard#compute}
	Compute interface{} `field:"required" json:"compute" yaml:"compute"`
	// The data source for event platform-based queries.
	//
	// Valid values are `logs`, `spans`, `network`, `rum`, `security_signals`, `profiles`, `audit`, `events`, `ci_tests`, `ci_pipelines`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/dashboard#group_by Dashboard#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// An array of index names to query in the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/dashboard#indexes Dashboard#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/dashboard#search Dashboard#search}
	Search *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinitionRequestQueryEventQuerySearch `field:"optional" json:"search" yaml:"search"`
	// Storage location (private beta).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.32.0/docs/resources/dashboard#storage Dashboard#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

