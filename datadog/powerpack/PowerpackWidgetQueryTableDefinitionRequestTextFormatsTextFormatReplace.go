// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryTableDefinitionRequestTextFormatsTextFormatReplace struct {
	// Table widget text format replace all type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/powerpack#type Powerpack#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Table Widget Match String.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/powerpack#with Powerpack#with}
	With *string `field:"required" json:"with" yaml:"with"`
	// Text that will be replaced. Must be used with type `substring`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.65.0/docs/resources/powerpack#substring Powerpack#substring}
	Substring *string `field:"optional" json:"substring" yaml:"substring"`
}

