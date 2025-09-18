// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryValueDefinitionRequestQueryProcessQuery struct {
	// The data source for process queries. Valid values are `process`, `container`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The process metric name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#metric Powerpack#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The aggregation methods available for metrics queries. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#aggregator Powerpack#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#cross_org_uuids Powerpack#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// Whether to normalize the CPU percentages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#is_normalized_cpu Powerpack#is_normalized_cpu}
	IsNormalizedCpu interface{} `field:"optional" json:"isNormalizedCpu" yaml:"isNormalizedCpu"`
	// The number of hits to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#limit Powerpack#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// The direction of the sort. Valid values are `asc`, `desc`. Defaults to `"desc"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#sort Powerpack#sort}
	Sort *string `field:"optional" json:"sort" yaml:"sort"`
	// An array of tags to filter by.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#tag_filters Powerpack#tag_filters}
	TagFilters *[]*string `field:"optional" json:"tagFilters" yaml:"tagFilters"`
	// The text to use as a filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/powerpack#text_filter Powerpack#text_filter}
	TextFilter *string `field:"optional" json:"textFilter" yaml:"textFilter"`
}

