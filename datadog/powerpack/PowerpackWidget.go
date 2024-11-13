// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidget struct {
	// alert_graph_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#alert_graph_definition Powerpack#alert_graph_definition}
	AlertGraphDefinition *PowerpackWidgetAlertGraphDefinition `field:"optional" json:"alertGraphDefinition" yaml:"alertGraphDefinition"`
	// alert_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#alert_value_definition Powerpack#alert_value_definition}
	AlertValueDefinition *PowerpackWidgetAlertValueDefinition `field:"optional" json:"alertValueDefinition" yaml:"alertValueDefinition"`
	// change_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#change_definition Powerpack#change_definition}
	ChangeDefinition *PowerpackWidgetChangeDefinition `field:"optional" json:"changeDefinition" yaml:"changeDefinition"`
	// check_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#check_status_definition Powerpack#check_status_definition}
	CheckStatusDefinition *PowerpackWidgetCheckStatusDefinition `field:"optional" json:"checkStatusDefinition" yaml:"checkStatusDefinition"`
	// distribution_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#distribution_definition Powerpack#distribution_definition}
	DistributionDefinition *PowerpackWidgetDistributionDefinition `field:"optional" json:"distributionDefinition" yaml:"distributionDefinition"`
	// event_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#event_stream_definition Powerpack#event_stream_definition}
	EventStreamDefinition *PowerpackWidgetEventStreamDefinition `field:"optional" json:"eventStreamDefinition" yaml:"eventStreamDefinition"`
	// event_timeline_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#event_timeline_definition Powerpack#event_timeline_definition}
	EventTimelineDefinition *PowerpackWidgetEventTimelineDefinition `field:"optional" json:"eventTimelineDefinition" yaml:"eventTimelineDefinition"`
	// free_text_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#free_text_definition Powerpack#free_text_definition}
	FreeTextDefinition *PowerpackWidgetFreeTextDefinition `field:"optional" json:"freeTextDefinition" yaml:"freeTextDefinition"`
	// geomap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#geomap_definition Powerpack#geomap_definition}
	GeomapDefinition *PowerpackWidgetGeomapDefinition `field:"optional" json:"geomapDefinition" yaml:"geomapDefinition"`
	// heatmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#heatmap_definition Powerpack#heatmap_definition}
	HeatmapDefinition *PowerpackWidgetHeatmapDefinition `field:"optional" json:"heatmapDefinition" yaml:"heatmapDefinition"`
	// hostmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#hostmap_definition Powerpack#hostmap_definition}
	HostmapDefinition *PowerpackWidgetHostmapDefinition `field:"optional" json:"hostmapDefinition" yaml:"hostmapDefinition"`
	// iframe_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#iframe_definition Powerpack#iframe_definition}
	IframeDefinition *PowerpackWidgetIframeDefinition `field:"optional" json:"iframeDefinition" yaml:"iframeDefinition"`
	// image_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#image_definition Powerpack#image_definition}
	ImageDefinition *PowerpackWidgetImageDefinition `field:"optional" json:"imageDefinition" yaml:"imageDefinition"`
	// list_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#list_stream_definition Powerpack#list_stream_definition}
	ListStreamDefinition *PowerpackWidgetListStreamDefinition `field:"optional" json:"listStreamDefinition" yaml:"listStreamDefinition"`
	// log_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#log_stream_definition Powerpack#log_stream_definition}
	LogStreamDefinition *PowerpackWidgetLogStreamDefinition `field:"optional" json:"logStreamDefinition" yaml:"logStreamDefinition"`
	// manage_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#manage_status_definition Powerpack#manage_status_definition}
	ManageStatusDefinition *PowerpackWidgetManageStatusDefinition `field:"optional" json:"manageStatusDefinition" yaml:"manageStatusDefinition"`
	// note_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#note_definition Powerpack#note_definition}
	NoteDefinition *PowerpackWidgetNoteDefinition `field:"optional" json:"noteDefinition" yaml:"noteDefinition"`
	// query_table_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#query_table_definition Powerpack#query_table_definition}
	QueryTableDefinition *PowerpackWidgetQueryTableDefinition `field:"optional" json:"queryTableDefinition" yaml:"queryTableDefinition"`
	// query_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#query_value_definition Powerpack#query_value_definition}
	QueryValueDefinition *PowerpackWidgetQueryValueDefinition `field:"optional" json:"queryValueDefinition" yaml:"queryValueDefinition"`
	// run_workflow_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#run_workflow_definition Powerpack#run_workflow_definition}
	RunWorkflowDefinition *PowerpackWidgetRunWorkflowDefinition `field:"optional" json:"runWorkflowDefinition" yaml:"runWorkflowDefinition"`
	// scatterplot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#scatterplot_definition Powerpack#scatterplot_definition}
	ScatterplotDefinition *PowerpackWidgetScatterplotDefinition `field:"optional" json:"scatterplotDefinition" yaml:"scatterplotDefinition"`
	// service_level_objective_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#service_level_objective_definition Powerpack#service_level_objective_definition}
	ServiceLevelObjectiveDefinition *PowerpackWidgetServiceLevelObjectiveDefinition `field:"optional" json:"serviceLevelObjectiveDefinition" yaml:"serviceLevelObjectiveDefinition"`
	// servicemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#servicemap_definition Powerpack#servicemap_definition}
	ServicemapDefinition *PowerpackWidgetServicemapDefinition `field:"optional" json:"servicemapDefinition" yaml:"servicemapDefinition"`
	// slo_list_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#slo_list_definition Powerpack#slo_list_definition}
	SloListDefinition *PowerpackWidgetSloListDefinition `field:"optional" json:"sloListDefinition" yaml:"sloListDefinition"`
	// sunburst_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#sunburst_definition Powerpack#sunburst_definition}
	SunburstDefinition *PowerpackWidgetSunburstDefinition `field:"optional" json:"sunburstDefinition" yaml:"sunburstDefinition"`
	// timeseries_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#timeseries_definition Powerpack#timeseries_definition}
	TimeseriesDefinition *PowerpackWidgetTimeseriesDefinition `field:"optional" json:"timeseriesDefinition" yaml:"timeseriesDefinition"`
	// toplist_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#toplist_definition Powerpack#toplist_definition}
	ToplistDefinition *PowerpackWidgetToplistDefinition `field:"optional" json:"toplistDefinition" yaml:"toplistDefinition"`
	// topology_map_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#topology_map_definition Powerpack#topology_map_definition}
	TopologyMapDefinition *PowerpackWidgetTopologyMapDefinition `field:"optional" json:"topologyMapDefinition" yaml:"topologyMapDefinition"`
	// trace_service_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#trace_service_definition Powerpack#trace_service_definition}
	TraceServiceDefinition *PowerpackWidgetTraceServiceDefinition `field:"optional" json:"traceServiceDefinition" yaml:"traceServiceDefinition"`
	// treemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#treemap_definition Powerpack#treemap_definition}
	TreemapDefinition *PowerpackWidgetTreemapDefinition `field:"optional" json:"treemapDefinition" yaml:"treemapDefinition"`
	// widget_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#widget_layout Powerpack#widget_layout}
	WidgetLayout *PowerpackWidgetWidgetLayout `field:"optional" json:"widgetLayout" yaml:"widgetLayout"`
}

