// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionRequestQueryEventQueryGroupBySort struct {
	// The aggregation methods for the event platform queries.
	//
	// Valid values are `count`, `cardinality`, `median`, `pc75`, `pc90`, `pc95`, `pc98`, `pc99`, `sum`, `min`, `max`, `avg`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/resources/dashboard#aggregation Dashboard#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The metric used for sorting group by results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/resources/dashboard#metric Dashboard#metric}
	Metric *string `field:"optional" json:"metric" yaml:"metric"`
	// Direction of sort. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/resources/dashboard#order Dashboard#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

