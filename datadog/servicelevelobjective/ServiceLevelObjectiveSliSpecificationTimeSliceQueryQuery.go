// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective


type ServiceLevelObjectiveSliSpecificationTimeSliceQueryQuery struct {
	// metric_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/service_level_objective#metric_query ServiceLevelObjective#metric_query}
	MetricQuery *ServiceLevelObjectiveSliSpecificationTimeSliceQueryQueryMetricQuery `field:"optional" json:"metricQuery" yaml:"metricQuery"`
}

