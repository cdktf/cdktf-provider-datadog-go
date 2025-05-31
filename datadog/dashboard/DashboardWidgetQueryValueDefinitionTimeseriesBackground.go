// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetQueryValueDefinitionTimeseriesBackground struct {
	// Whether the Timeseries is made using an area or bars. Valid values are `bars`, `area`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/dashboard#yaxis Dashboard#yaxis}
	Yaxis *DashboardWidgetQueryValueDefinitionTimeseriesBackgroundYaxis `field:"optional" json:"yaxis" yaml:"yaxis"`
}

