// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetListStreamDefinitionRequestColumns struct {
	// Widget column field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#field Powerpack#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// Widget column width. Valid values are `auto`, `compact`, `full`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#width Powerpack#width}
	Width *string `field:"optional" json:"width" yaml:"width"`
}

