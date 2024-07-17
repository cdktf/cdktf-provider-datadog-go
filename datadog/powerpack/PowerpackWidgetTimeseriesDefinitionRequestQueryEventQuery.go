// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuery struct {
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#compute Powerpack#compute}
	Compute interface{} `field:"required" json:"compute" yaml:"compute"`
	// The data source for event platform-based queries.
	//
	// Valid values are `logs`, `spans`, `network`, `rum`, `security_signals`, `profiles`, `audit`, `events`, `ci_tests`, `ci_pipelines`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#data_source Powerpack#data_source}
	DataSource *string `field:"required" json:"dataSource" yaml:"dataSource"`
	// The name of query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source organization UUID for cross organization queries. Feature in Private Beta.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#cross_org_uuids Powerpack#cross_org_uuids}
	CrossOrgUuids *[]*string `field:"optional" json:"crossOrgUuids" yaml:"crossOrgUuids"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#group_by Powerpack#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// An array of index names to query in the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#indexes Powerpack#indexes}
	Indexes *[]*string `field:"optional" json:"indexes" yaml:"indexes"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#search Powerpack#search}
	Search *PowerpackWidgetTimeseriesDefinitionRequestQueryEventQuerySearch `field:"optional" json:"search" yaml:"search"`
	// Storage location (private beta).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.41.0/docs/resources/powerpack#storage Powerpack#storage}
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
}

