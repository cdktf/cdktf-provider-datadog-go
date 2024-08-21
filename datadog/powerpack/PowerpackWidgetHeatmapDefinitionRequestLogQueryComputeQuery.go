// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetHeatmapDefinitionRequestLogQueryComputeQuery struct {
	// The aggregation method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/powerpack#aggregation Powerpack#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/powerpack#facet Powerpack#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// Define the time interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/powerpack#interval Powerpack#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

