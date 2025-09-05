// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetScatterplotDefinitionRequestScatterplotTableQueryProcessQuery struct {
	// The data source for process queries. Valid values are `process`, `container`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#data_source Dashboard#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The process metric name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#metric Dashboard#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The aggregation methods available for metrics queries. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#aggregator Dashboard#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#cross_org_uuids Dashboard#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Whether to normalize the CPU percentages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#is_normalized_cpu Dashboard#is_normalized_cpu}
	IsNormalizedCpu interface{} `field:"optional" json:"isNormalizedCpu" yaml:"isNormalizedCpu"`
	// The number of hits to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#limit Dashboard#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The direction of the sort. Valid values are `asc`, `desc`. Defaults to `"desc"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#sort Dashboard#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// An array of tags to filter by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#tag_filters Dashboard#tag_filters}
	TagFilters *[]*string `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// The text to use as a filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/dashboard#text_filter Dashboard#text_filter}
	TextFilter *string `field:"optional" json:"textFilter" yaml:"textFilter"`
}

