// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSplitGraphDefinitionSourceWidgetDefinitionQueryTableDefinitionRequestQueryApmDependencyStatsQuery struct {
	// The data source for APM Dependency Stats queries. Valid values are `apm_dependency_stats`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// APM environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#env Dashboard#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of operation on service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#operation_name Dashboard#operation_name}
	OperationName *string `field:"required" json:"operationName" yaml:"operationName"`
	// APM resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#resource_name Dashboard#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// APM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#service Dashboard#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// APM statistic. Valid values are `avg_duration`, `avg_root_duration`, `avg_spans_per_trace`, `error_rate`, `pct_exec_time`, `pct_of_traces`, `total_traces_count`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#stat Dashboard#stat}
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#cross_org_uuids Dashboard#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Determines whether stats for upstream or downstream dependencies should be queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#is_upstream Dashboard#is_upstream}
	IsUpstream interface{} `field:"optional" json:"isUpstream" yaml:"isUpstream"`
	// The name of the second primary tag used within APM; required when `primary_tag_value` is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#primary_tag_name Dashboard#primary_tag_name}
	PrimaryTagName *string `field:"optional" json:"primaryTagName" yaml:"primaryTagName"`
	// Filter APM data by the second primary tag. `primary_tag_name` must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#primary_tag_value Dashboard#primary_tag_value}
	PrimaryTagValue *string `field:"optional" json:"primaryTagValue" yaml:"primaryTagValue"`
}

