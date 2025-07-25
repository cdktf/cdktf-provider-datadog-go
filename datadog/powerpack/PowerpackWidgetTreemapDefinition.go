// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTreemapDefinition struct {
	// custom_links block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#custom_links Powerpack#custom_links}
	CustomLinks interface{} `field:"optional" json:"customLinks" yaml:"customLinks"`
	// request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#request Powerpack#request}
	Request interface{} `field:"optional" json:"request" yaml:"request"`
	// The title of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/powerpack#title Powerpack#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

