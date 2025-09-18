// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinition struct {
	// change_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#change_definition Dashboard#change_definition}
	ChangeDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionChangeDefinition `field:"optional" json:"changeDefinition" yaml:"changeDefinition"`
	// geomap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#geomap_definition Dashboard#geomap_definition}
	GeomapDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionGeomapDefinition `field:"optional" json:"geomapDefinition" yaml:"geomapDefinition"`
	// query_table_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#query_table_definition Dashboard#query_table_definition}
	QueryTableDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinition `field:"optional" json:"queryTableDefinition" yaml:"queryTableDefinition"`
	// query_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#query_value_definition Dashboard#query_value_definition}
	QueryValueDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryValueDefinition `field:"optional" json:"queryValueDefinition" yaml:"queryValueDefinition"`
	// scatterplot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#scatterplot_definition Dashboard#scatterplot_definition}
	ScatterplotDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionScatterplotDefinition `field:"optional" json:"scatterplotDefinition" yaml:"scatterplotDefinition"`
	// sunburst_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#sunburst_definition Dashboard#sunburst_definition}
	SunburstDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionSunburstDefinition `field:"optional" json:"sunburstDefinition" yaml:"sunburstDefinition"`
	// timeseries_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#timeseries_definition Dashboard#timeseries_definition}
	TimeseriesDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinition `field:"optional" json:"timeseriesDefinition" yaml:"timeseriesDefinition"`
	// toplist_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#toplist_definition Dashboard#toplist_definition}
	ToplistDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionToplistDefinition `field:"optional" json:"toplistDefinition" yaml:"toplistDefinition"`
	// treemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/dashboard#treemap_definition Dashboard#treemap_definition}
	TreemapDefinition *DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTreemapDefinition `field:"optional" json:"treemapDefinition" yaml:"treemapDefinition"`
}

