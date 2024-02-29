// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetHostmapDefinitionRequest struct {
	// fill block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/powerpack#fill Powerpack#fill}
	Fill interface{} `field:"optional" json:"fill" yaml:"fill"`
	// size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/powerpack#size Powerpack#size}
	Size interface{} `field:"optional" json:"size" yaml:"size"`
}

