// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetSunburstDefinitionLegendTable struct {
	// The type of legend (table or none). Valid values are `table`, `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.1/docs/resources/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

