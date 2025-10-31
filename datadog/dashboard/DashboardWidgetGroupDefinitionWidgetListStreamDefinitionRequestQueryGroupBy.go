// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetListStreamDefinitionRequestQueryGroupBy struct {
	// Facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/dashboard#facet Dashboard#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
}

