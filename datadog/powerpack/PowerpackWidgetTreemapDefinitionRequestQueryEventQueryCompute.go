// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTreemapDefinitionRequestQueryEventQueryCompute struct {
	// The aggregation methods for event platform queries.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/powerpack#aggregation Powerpack#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// A time interval in milliseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/powerpack#interval Powerpack#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The measurable attribute to compute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/powerpack#metric Powerpack#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
}

