package dashboard


type DashboardWidget struct {
	// alert_graph_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#alert_graph_definition Dashboard#alert_graph_definition}
	AlertGraphDefinition *DashboardWidgetAlertGraphDefinition `field:"optional" json:"alertGraphDefinition" yaml:"alertGraphDefinition"`
	// alert_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#alert_value_definition Dashboard#alert_value_definition}
	AlertValueDefinition *DashboardWidgetAlertValueDefinition `field:"optional" json:"alertValueDefinition" yaml:"alertValueDefinition"`
	// change_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#change_definition Dashboard#change_definition}
	ChangeDefinition *DashboardWidgetChangeDefinition `field:"optional" json:"changeDefinition" yaml:"changeDefinition"`
	// check_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#check_status_definition Dashboard#check_status_definition}
	CheckStatusDefinition *DashboardWidgetCheckStatusDefinition `field:"optional" json:"checkStatusDefinition" yaml:"checkStatusDefinition"`
	// distribution_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#distribution_definition Dashboard#distribution_definition}
	DistributionDefinition *DashboardWidgetDistributionDefinition `field:"optional" json:"distributionDefinition" yaml:"distributionDefinition"`
	// event_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#event_stream_definition Dashboard#event_stream_definition}
	EventStreamDefinition *DashboardWidgetEventStreamDefinition `field:"optional" json:"eventStreamDefinition" yaml:"eventStreamDefinition"`
	// event_timeline_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#event_timeline_definition Dashboard#event_timeline_definition}
	EventTimelineDefinition *DashboardWidgetEventTimelineDefinition `field:"optional" json:"eventTimelineDefinition" yaml:"eventTimelineDefinition"`
	// free_text_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#free_text_definition Dashboard#free_text_definition}
	FreeTextDefinition *DashboardWidgetFreeTextDefinition `field:"optional" json:"freeTextDefinition" yaml:"freeTextDefinition"`
	// geomap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#geomap_definition Dashboard#geomap_definition}
	GeomapDefinition *DashboardWidgetGeomapDefinition `field:"optional" json:"geomapDefinition" yaml:"geomapDefinition"`
	// group_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#group_definition Dashboard#group_definition}
	GroupDefinition *DashboardWidgetGroupDefinition `field:"optional" json:"groupDefinition" yaml:"groupDefinition"`
	// heatmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#heatmap_definition Dashboard#heatmap_definition}
	HeatmapDefinition *DashboardWidgetHeatmapDefinition `field:"optional" json:"heatmapDefinition" yaml:"heatmapDefinition"`
	// hostmap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#hostmap_definition Dashboard#hostmap_definition}
	HostmapDefinition *DashboardWidgetHostmapDefinition `field:"optional" json:"hostmapDefinition" yaml:"hostmapDefinition"`
	// iframe_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#iframe_definition Dashboard#iframe_definition}
	IframeDefinition *DashboardWidgetIframeDefinition `field:"optional" json:"iframeDefinition" yaml:"iframeDefinition"`
	// image_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#image_definition Dashboard#image_definition}
	ImageDefinition *DashboardWidgetImageDefinition `field:"optional" json:"imageDefinition" yaml:"imageDefinition"`
	// list_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#list_stream_definition Dashboard#list_stream_definition}
	ListStreamDefinition *DashboardWidgetListStreamDefinition `field:"optional" json:"listStreamDefinition" yaml:"listStreamDefinition"`
	// log_stream_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#log_stream_definition Dashboard#log_stream_definition}
	LogStreamDefinition *DashboardWidgetLogStreamDefinition `field:"optional" json:"logStreamDefinition" yaml:"logStreamDefinition"`
	// manage_status_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#manage_status_definition Dashboard#manage_status_definition}
	ManageStatusDefinition *DashboardWidgetManageStatusDefinition `field:"optional" json:"manageStatusDefinition" yaml:"manageStatusDefinition"`
	// note_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#note_definition Dashboard#note_definition}
	NoteDefinition *DashboardWidgetNoteDefinition `field:"optional" json:"noteDefinition" yaml:"noteDefinition"`
	// query_table_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#query_table_definition Dashboard#query_table_definition}
	QueryTableDefinition *DashboardWidgetQueryTableDefinition `field:"optional" json:"queryTableDefinition" yaml:"queryTableDefinition"`
	// query_value_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#query_value_definition Dashboard#query_value_definition}
	QueryValueDefinition *DashboardWidgetQueryValueDefinition `field:"optional" json:"queryValueDefinition" yaml:"queryValueDefinition"`
	// run_workflow_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#run_workflow_definition Dashboard#run_workflow_definition}
	RunWorkflowDefinition *DashboardWidgetRunWorkflowDefinition `field:"optional" json:"runWorkflowDefinition" yaml:"runWorkflowDefinition"`
	// scatterplot_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#scatterplot_definition Dashboard#scatterplot_definition}
	ScatterplotDefinition *DashboardWidgetScatterplotDefinition `field:"optional" json:"scatterplotDefinition" yaml:"scatterplotDefinition"`
	// service_level_objective_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#service_level_objective_definition Dashboard#service_level_objective_definition}
	ServiceLevelObjectiveDefinition *DashboardWidgetServiceLevelObjectiveDefinition `field:"optional" json:"serviceLevelObjectiveDefinition" yaml:"serviceLevelObjectiveDefinition"`
	// servicemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#servicemap_definition Dashboard#servicemap_definition}
	ServicemapDefinition *DashboardWidgetServicemapDefinition `field:"optional" json:"servicemapDefinition" yaml:"servicemapDefinition"`
	// slo_list_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#slo_list_definition Dashboard#slo_list_definition}
	SloListDefinition *DashboardWidgetSloListDefinition `field:"optional" json:"sloListDefinition" yaml:"sloListDefinition"`
	// sunburst_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#sunburst_definition Dashboard#sunburst_definition}
	SunburstDefinition *DashboardWidgetSunburstDefinition `field:"optional" json:"sunburstDefinition" yaml:"sunburstDefinition"`
	// timeseries_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#timeseries_definition Dashboard#timeseries_definition}
	TimeseriesDefinition *DashboardWidgetTimeseriesDefinition `field:"optional" json:"timeseriesDefinition" yaml:"timeseriesDefinition"`
	// toplist_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#toplist_definition Dashboard#toplist_definition}
	ToplistDefinition *DashboardWidgetToplistDefinition `field:"optional" json:"toplistDefinition" yaml:"toplistDefinition"`
	// topology_map_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#topology_map_definition Dashboard#topology_map_definition}
	TopologyMapDefinition *DashboardWidgetTopologyMapDefinition `field:"optional" json:"topologyMapDefinition" yaml:"topologyMapDefinition"`
	// trace_service_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#trace_service_definition Dashboard#trace_service_definition}
	TraceServiceDefinition *DashboardWidgetTraceServiceDefinition `field:"optional" json:"traceServiceDefinition" yaml:"traceServiceDefinition"`
	// treemap_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#treemap_definition Dashboard#treemap_definition}
	TreemapDefinition *DashboardWidgetTreemapDefinition `field:"optional" json:"treemapDefinition" yaml:"treemapDefinition"`
	// widget_layout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.24.0/docs/resources/dashboard#widget_layout Dashboard#widget_layout}
	WidgetLayout *DashboardWidgetWidgetLayout `field:"optional" json:"widgetLayout" yaml:"widgetLayout"`
}

