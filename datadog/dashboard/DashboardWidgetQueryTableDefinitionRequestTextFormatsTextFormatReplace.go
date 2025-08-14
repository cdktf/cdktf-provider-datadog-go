// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace struct {
	// Table widget text format replace all type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Table Widget Match String.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#with Dashboard#with}
	With *string `field:"required" json:"with" yaml:"with"`
	// Text that will be replaced. Must be used with type `substring`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#substring Dashboard#substring}
	Substring *string `field:"optional" json:"substring" yaml:"substring"`
}

