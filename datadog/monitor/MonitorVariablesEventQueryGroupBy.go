// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitor


type MonitorVariablesEventQueryGroupBy struct {
	// The event facet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/monitor#facet Monitor#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
	// The number of groups to return.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/monitor#limit Monitor#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/monitor#sort Monitor#sort}
	Sort *MonitorVariablesEventQueryGroupBySort `field:"optional" json:"sort" yaml:"sort"`
}

