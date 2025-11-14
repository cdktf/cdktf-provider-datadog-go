// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetToplistDefinitionStyle struct {
	// display block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/powerpack#display Powerpack#display}
	Display interface{} `field:"optional" json:"display" yaml:"display"`
	// The color palette for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/powerpack#palette Powerpack#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
	// The scaling mode for the widget. Valid values are `absolute`, `relative`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/powerpack#scaling Powerpack#scaling}
	Scaling *string `field:"optional" json:"scaling" yaml:"scaling"`
}

