// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetSunburstDefinitionRequestQuerySloQuery struct {
	// The data source for SLO queries. Valid values are `slo`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// SLO measures queries. Valid values are `good_events`, `bad_events`, `good_minutes`, `bad_minutes`, `slo_status`, `error_budget_remaining`, `burn_rate`, `error_budget_burndown`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#measure Powerpack#measure}
	Measure *string `field:"required" json:"measure" yaml:"measure"`
	// ID of an SLO to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#slo_id Powerpack#slo_id}
	SloId *string `field:"required" json:"sloId" yaml:"sloId"`
	// Additional filters applied to the SLO query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#additional_query_filters Powerpack#additional_query_filters}
	AdditionalQueryFilters *string `field:"optional" json:"additionalQueryFilters" yaml:"additionalQueryFilters"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#cross_org_uuids Powerpack#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Group mode to query measures. Valid values are `overall`, `components`. Defaults to `"overall"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#group_mode Powerpack#group_mode}
	GroupMode *string `field:"optional" json:"groupMode" yaml:"groupMode"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// type of the SLO to query. Valid values are `metric`, `monitor`, `time_slice`. Defaults to `"metric"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#slo_query_type Powerpack#slo_query_type}
	SloQueryType *string `field:"optional" json:"sloQueryType" yaml:"sloQueryType"`
}

