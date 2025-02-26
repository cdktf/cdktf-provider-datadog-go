// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetScatterplotDefinitionRequestScatterplotTableQueryApmResourceStatsQuery struct {
	// The data source for APM Resource Stats queries. Valid values are `apm_resource_stats`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// APM environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#env Powerpack#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// APM service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#service Powerpack#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// APM statistic. Valid values are `errors`, `error_rate`, `hits`, `latency_avg`, `latency_distribution`, `latency_max`, `latency_p50`, `latency_p75`, `latency_p90`, `latency_p95`, `latency_p99`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#stat Powerpack#stat}
	Stat *string `field:"required" json:"stat" yaml:"stat"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#cross_org_uuids Powerpack#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Array of fields to group results by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#group_by Powerpack#group_by}
	GroupBy *[]*string `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Name of operation on service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#operation_name Powerpack#operation_name}
	OperationName *string `field:"optional" json:"operationName" yaml:"operationName"`
	// The name of the second primary tag used within APM; required when `primary_tag_value` is specified. See https://docs.datadoghq.com/tracing/guide/setting_primary_tags_to_scope/#add-a-second-primary-tag-in-datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#primary_tag_name Powerpack#primary_tag_name}
	PrimaryTagName *string `field:"optional" json:"primaryTagName" yaml:"primaryTagName"`
	// Filter APM data by the second primary tag. `primary_tag_name` must also be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#primary_tag_value Powerpack#primary_tag_value}
	PrimaryTagValue *string `field:"optional" json:"primaryTagValue" yaml:"primaryTagValue"`
	// APM resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.56.0/docs/resources/powerpack#resource_name Powerpack#resource_name}
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
}

