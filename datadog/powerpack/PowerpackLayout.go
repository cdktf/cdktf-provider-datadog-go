// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackLayout struct {
	// The height of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#height Powerpack#height}
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The width of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#width Powerpack#width}
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// The position of the widget on the x (horizontal) axis. Should be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#x Powerpack#x}
	X *float64 `field:"optional" json:"x" yaml:"x"`
	// The position of the widget on the y (vertical) axis. Should be greater than or equal to 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#y Powerpack#y}
	Y *float64 `field:"optional" json:"y" yaml:"y"`
}

