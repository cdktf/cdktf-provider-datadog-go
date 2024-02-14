// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetListStreamDefinitionRequestColumns struct {
	// Widget column field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/dashboard#field Dashboard#field}
	Field *string `field:"required" json:"field" yaml:"field"`
	// Widget column width. Valid values are `auto`, `compact`, `full`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/dashboard#width Dashboard#width}
	Width *string `field:"required" json:"width" yaml:"width"`
}

