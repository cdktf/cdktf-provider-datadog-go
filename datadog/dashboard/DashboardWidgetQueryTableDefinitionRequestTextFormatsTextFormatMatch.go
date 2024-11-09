// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch struct {
	// Match or compare option. Valid values are `is`, `is_not`, `contains`, `does_not_contain`, `starts_with`, `ends_with`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.47.0/docs/resources/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Table Widget Match String.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.47.0/docs/resources/dashboard#value Dashboard#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

