// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetToplistDefinitionRequestQueryCloudCostQuery struct {
	// The data source for cloud cost queries. Valid values are `cloud_cost`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The cloud cost query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/powerpack#query Powerpack#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The aggregation methods available for cloud cost queries. Valid values are `avg`, `last`, `max`, `min`, `sum`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/powerpack#aggregator Powerpack#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
}

