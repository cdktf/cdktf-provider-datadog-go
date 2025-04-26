// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecificationTimeSliceQueryQueryMetricQuery struct {
	// The name of the query for use in formulas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/service_level_objective#name ServiceLevelObjective#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The metrics query definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/service_level_objective#query ServiceLevelObjective#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// The data source for metrics queries. Defaults to `"metrics"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/service_level_objective#data_source ServiceLevelObjective#data_source}
	DataSource *string `field:"optional" json:"dataSource" yaml:"dataSource"`
}

