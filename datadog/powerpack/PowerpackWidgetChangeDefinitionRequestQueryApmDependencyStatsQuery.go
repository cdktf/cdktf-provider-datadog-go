// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetChangeDefinitionRequestQueryApmDependencyStatsQuery struct {
	// The data source for APM Dependency Stats queries. Valid values are `apm_dependency_stats`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// APM environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#env Powerpack#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of operation on service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#operation_name Powerpack#operation_name}
	OperationName *string `field:"required" json:"operationName" yaml:"operationName"`
	// APM resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#resource_name Powerpack#resource_name}
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// APM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#service Powerpack#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// APM statistic. Valid values are `avg_duration`, `avg_root_duration`, `avg_spans_per_trace`, `error_rate`, `pct_exec_time`, `pct_of_traces`, `total_traces_count`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#stat Powerpack#stat}
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// Determines whether stats for upstream or downstream dependencies should be queried.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#is_upstream Powerpack#is_upstream}
	IsUpstream interface{} `field:"optional" json:"isUpstream" yaml:"isUpstream"`
	// The name of the second primary tag used within APM; required when `primary_tag_value` is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#primary_tag_name Powerpack#primary_tag_name}
	PrimaryTagName *string `field:"optional" json:"primaryTagName" yaml:"primaryTagName"`
	// Filter APM data by the second primary tag. `primary_tag_name` must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.34.0/docs/resources/powerpack#primary_tag_value Powerpack#primary_tag_value}
	PrimaryTagValue *string `field:"optional" json:"primaryTagValue" yaml:"primaryTagValue"`
}

