// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecificationTimeSlice struct {
	// The comparator used to compare the SLI value to the threshold. Valid values are `>`, `>=`, `<`, `<=`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.47.0/docs/resources/service_level_objective#comparator ServiceLevelObjective#comparator}
	Comparator *string `field:"required" json:"comparator" yaml:"comparator"`
	// query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.47.0/docs/resources/service_level_objective#query ServiceLevelObjective#query}
	Query *ServiceLevelObjectiveSliSpecificationTimeSliceQuery `field:"required" json:"query" yaml:"query"`
	// The threshold value to which each SLI value will be compared.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.47.0/docs/resources/service_level_objective#threshold ServiceLevelObjective#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The interval used when querying data, which defines the size of a time slice.
	//
	// Valid values are `60`, `300`. Defaults to `300`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.47.0/docs/resources/service_level_objective#query_interval_seconds ServiceLevelObjective#query_interval_seconds}
	QueryIntervalSeconds *float64 `field:"optional" json:"queryIntervalSeconds" yaml:"queryIntervalSeconds"`
}

