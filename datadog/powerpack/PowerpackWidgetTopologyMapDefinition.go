// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTopologyMapDefinition struct {
	// custom_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#custom_link Powerpack#custom_link}
	CustomLink interface{} `field:"optional" json:"customLink" yaml:"customLink"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#request Powerpack#request}
	Request interface{} `field:"optional" json:"request" yaml:"request"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#title Powerpack#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The alignment of the widget's title. Valid values are `center`, `left`, `right`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#title_align Powerpack#title_align}
	TitleAlign *string `field:"optional" json:"titleAlign" yaml:"titleAlign"`
	// The size of the widget's title (defaults to 16).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#title_size Powerpack#title_size}
	TitleSize *string `field:"optional" json:"titleSize" yaml:"titleSize"`
}

