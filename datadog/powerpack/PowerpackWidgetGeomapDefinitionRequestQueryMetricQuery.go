// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetGeomapDefinitionRequestQueryMetricQuery struct {
	// The name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metrics query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#query Powerpack#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The aggregation methods available for metrics queries. Valid values are `avg`, `min`, `max`, `sum`, `last`, `area`, `l2norm`, `percentile`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#aggregator Powerpack#aggregator}
	Aggregator *string `field:"optional" json:"aggregator" yaml:"aggregator"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#cross_org_uuids Powerpack#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// The data source for metrics queries. Defaults to `"metrics"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
}

