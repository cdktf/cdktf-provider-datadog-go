// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatMatch struct {
	// Match or compare option. Valid values are `is`, `is_not`, `contains`, `does_not_contain`, `starts_with`, `ends_with`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#type Powerpack#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Table Widget Match String.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/powerpack#value Powerpack#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

