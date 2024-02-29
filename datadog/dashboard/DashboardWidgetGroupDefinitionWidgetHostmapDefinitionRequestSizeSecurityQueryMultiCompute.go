// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetHostmapDefinitionRequestSizeSecurityQueryMultiCompute struct {
	// The aggregation method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#aggregation Dashboard#aggregation}
	Aggregation *string `field:"required" json:"aggregation" yaml:"aggregation"`
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#facet Dashboard#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// Define the time interval in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/dashboard#interval Dashboard#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
}

