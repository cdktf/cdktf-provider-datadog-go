// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetHostmapDefinitionRequest struct {
	// fill block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/dashboard#fill Dashboard#fill}
	Fill interface{} `field:"optional" json:"fill" yaml:"fill"`
	// size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/dashboard#size Dashboard#size}
	Size interface{} `field:"optional" json:"size" yaml:"size"`
}

