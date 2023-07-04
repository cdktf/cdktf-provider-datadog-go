package dashboard


type DashboardWidgetSunburstDefinitionRequestQuerySloQuery struct {
	// The data source for slo queries. Valid values are `slo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// SLO measures queries. Valid values are `good_events`, `bad_events`, `slo_status`, `error_budget_remaining`, `burn_rate`, `error_budget_burndown`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#measure Dashboard#measure}
	Measure *string `field:"required" json:"measure" yaml:"measure"`
	// ID of an SLO to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#slo_id Dashboard#slo_id}
	SloId *string `field:"required" json:"sloId" yaml:"sloId"`
	// Group mode to query measures. Valid values are `overall`, `components`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#group_mode Dashboard#group_mode}
	GroupMode *string `field:"optional" json:"groupMode" yaml:"groupMode"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// type of the SLO to query. Valid values are `metric`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.27.0/docs/resources/dashboard#slo_query_type Dashboard#slo_query_type}
	SloQueryType *string `field:"optional" json:"sloQueryType" yaml:"sloQueryType"`
}

