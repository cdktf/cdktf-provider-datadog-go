// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetDistributionDefinitionRequestApmStatsQueryColumns struct {
	// The column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A user-assigned alias for the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/dashboard#alias Dashboard#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// A list of display modes for each table cell. Valid values are `number`, `bar`, `trend`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/dashboard#cell_display_mode Dashboard#cell_display_mode}
	CellDisplayMode *string `field:"optional" json:"cellDisplayMode" yaml:"cellDisplayMode"`
	// Widget sorting methods. Valid values are `asc`, `desc`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/dashboard#order Dashboard#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

