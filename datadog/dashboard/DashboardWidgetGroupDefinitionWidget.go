// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidget struct {
	// alert_graph_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#alert_graph_definition Dashboard#alert_graph_definition}
	AlertGraphDefinition *DashboardWidgetGroupDefinitionWidgetAlertGraphDefinition `field:"optional" json:"alertGraphDefinition" yaml:"alertGraphDefinition"`
	// alert_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#alert_value_definition Dashboard#alert_value_definition}
	AlertValueDefinition *DashboardWidgetGroupDefinitionWidgetAlertValueDefinition `field:"optional" json:"alertValueDefinition" yaml:"alertValueDefinition"`
	// change_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#change_definition Dashboard#change_definition}
	ChangeDefinition *DashboardWidgetGroupDefinitionWidgetChangeDefinition `field:"optional" json:"changeDefinition" yaml:"changeDefinition"`
	// check_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#check_status_definition Dashboard#check_status_definition}
	CheckStatusDefinition *DashboardWidgetGroupDefinitionWidgetCheckStatusDefinition `field:"optional" json:"checkStatusDefinition" yaml:"checkStatusDefinition"`
	// distribution_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#distribution_definition Dashboard#distribution_definition}
	DistributionDefinition *DashboardWidgetGroupDefinitionWidgetDistributionDefinition `field:"optional" json:"distributionDefinition" yaml:"distributionDefinition"`
	// event_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#event_stream_definition Dashboard#event_stream_definition}
	EventStreamDefinition *DashboardWidgetGroupDefinitionWidgetEventStreamDefinition `field:"optional" json:"eventStreamDefinition" yaml:"eventStreamDefinition"`
	// event_timeline_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#event_timeline_definition Dashboard#event_timeline_definition}
	EventTimelineDefinition *DashboardWidgetGroupDefinitionWidgetEventTimelineDefinition `field:"optional" json:"eventTimelineDefinition" yaml:"eventTimelineDefinition"`
	// free_text_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#free_text_definition Dashboard#free_text_definition}
	FreeTextDefinition *DashboardWidgetGroupDefinitionWidgetFreeTextDefinition `field:"optional" json:"freeTextDefinition" yaml:"freeTextDefinition"`
	// geomap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#geomap_definition Dashboard#geomap_definition}
	GeomapDefinition *DashboardWidgetGroupDefinitionWidgetGeomapDefinition `field:"optional" json:"geomapDefinition" yaml:"geomapDefinition"`
	// heatmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#heatmap_definition Dashboard#heatmap_definition}
	HeatmapDefinition *DashboardWidgetGroupDefinitionWidgetHeatmapDefinition `field:"optional" json:"heatmapDefinition" yaml:"heatmapDefinition"`
	// hostmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#hostmap_definition Dashboard#hostmap_definition}
	HostmapDefinition *DashboardWidgetGroupDefinitionWidgetHostmapDefinition `field:"optional" json:"hostmapDefinition" yaml:"hostmapDefinition"`
	// iframe_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#iframe_definition Dashboard#iframe_definition}
	IframeDefinition *DashboardWidgetGroupDefinitionWidgetIframeDefinition `field:"optional" json:"iframeDefinition" yaml:"iframeDefinition"`
	// image_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#image_definition Dashboard#image_definition}
	ImageDefinition *DashboardWidgetGroupDefinitionWidgetImageDefinition `field:"optional" json:"imageDefinition" yaml:"imageDefinition"`
	// list_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#list_stream_definition Dashboard#list_stream_definition}
	ListStreamDefinition *DashboardWidgetGroupDefinitionWidgetListStreamDefinition `field:"optional" json:"listStreamDefinition" yaml:"listStreamDefinition"`
	// log_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#log_stream_definition Dashboard#log_stream_definition}
	LogStreamDefinition *DashboardWidgetGroupDefinitionWidgetLogStreamDefinition `field:"optional" json:"logStreamDefinition" yaml:"logStreamDefinition"`
	// manage_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#manage_status_definition Dashboard#manage_status_definition}
	ManageStatusDefinition *DashboardWidgetGroupDefinitionWidgetManageStatusDefinition `field:"optional" json:"manageStatusDefinition" yaml:"manageStatusDefinition"`
	// note_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#note_definition Dashboard#note_definition}
	NoteDefinition *DashboardWidgetGroupDefinitionWidgetNoteDefinition `field:"optional" json:"noteDefinition" yaml:"noteDefinition"`
	// powerpack_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#powerpack_definition Dashboard#powerpack_definition}
	PowerpackDefinition *DashboardWidgetGroupDefinitionWidgetPowerpackDefinition `field:"optional" json:"powerpackDefinition" yaml:"powerpackDefinition"`
	// query_table_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#query_table_definition Dashboard#query_table_definition}
	QueryTableDefinition *DashboardWidgetGroupDefinitionWidgetQueryTableDefinition `field:"optional" json:"queryTableDefinition" yaml:"queryTableDefinition"`
	// query_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#query_value_definition Dashboard#query_value_definition}
	QueryValueDefinition *DashboardWidgetGroupDefinitionWidgetQueryValueDefinition `field:"optional" json:"queryValueDefinition" yaml:"queryValueDefinition"`
	// run_workflow_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#run_workflow_definition Dashboard#run_workflow_definition}
	RunWorkflowDefinition *DashboardWidgetGroupDefinitionWidgetRunWorkflowDefinition `field:"optional" json:"runWorkflowDefinition" yaml:"runWorkflowDefinition"`
	// scatterplot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#scatterplot_definition Dashboard#scatterplot_definition}
	ScatterplotDefinition *DashboardWidgetGroupDefinitionWidgetScatterplotDefinition `field:"optional" json:"scatterplotDefinition" yaml:"scatterplotDefinition"`
	// service_level_objective_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#service_level_objective_definition Dashboard#service_level_objective_definition}
	ServiceLevelObjectiveDefinition *DashboardWidgetGroupDefinitionWidgetServiceLevelObjectiveDefinition `field:"optional" json:"serviceLevelObjectiveDefinition" yaml:"serviceLevelObjectiveDefinition"`
	// servicemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#servicemap_definition Dashboard#servicemap_definition}
	ServicemapDefinition *DashboardWidgetGroupDefinitionWidgetServicemapDefinition `field:"optional" json:"servicemapDefinition" yaml:"servicemapDefinition"`
	// slo_list_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#slo_list_definition Dashboard#slo_list_definition}
	SloListDefinition *DashboardWidgetGroupDefinitionWidgetSloListDefinition `field:"optional" json:"sloListDefinition" yaml:"sloListDefinition"`
	// split_graph_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#split_graph_definition Dashboard#split_graph_definition}
	SplitGraphDefinition *DashboardWidgetGroupDefinitionWidgetSplitGraphDefinition `field:"optional" json:"splitGraphDefinition" yaml:"splitGraphDefinition"`
	// sunburst_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#sunburst_definition Dashboard#sunburst_definition}
	SunburstDefinition *DashboardWidgetGroupDefinitionWidgetSunburstDefinition `field:"optional" json:"sunburstDefinition" yaml:"sunburstDefinition"`
	// timeseries_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#timeseries_definition Dashboard#timeseries_definition}
	TimeseriesDefinition *DashboardWidgetGroupDefinitionWidgetTimeseriesDefinition `field:"optional" json:"timeseriesDefinition" yaml:"timeseriesDefinition"`
	// toplist_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#toplist_definition Dashboard#toplist_definition}
	ToplistDefinition *DashboardWidgetGroupDefinitionWidgetToplistDefinition `field:"optional" json:"toplistDefinition" yaml:"toplistDefinition"`
	// topology_map_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#topology_map_definition Dashboard#topology_map_definition}
	TopologyMapDefinition *DashboardWidgetGroupDefinitionWidgetTopologyMapDefinition `field:"optional" json:"topologyMapDefinition" yaml:"topologyMapDefinition"`
	// trace_service_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#trace_service_definition Dashboard#trace_service_definition}
	TraceServiceDefinition *DashboardWidgetGroupDefinitionWidgetTraceServiceDefinition `field:"optional" json:"traceServiceDefinition" yaml:"traceServiceDefinition"`
	// treemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#treemap_definition Dashboard#treemap_definition}
	TreemapDefinition *DashboardWidgetGroupDefinitionWidgetTreemapDefinition `field:"optional" json:"treemapDefinition" yaml:"treemapDefinition"`
	// widget_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/dashboard#widget_layout Dashboard#widget_layout}
	WidgetLayout *DashboardWidgetGroupDefinitionWidgetWidgetLayout `field:"optional" json:"widgetLayout" yaml:"widgetLayout"`
}

