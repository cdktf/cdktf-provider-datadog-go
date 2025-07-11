// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetToplistDefinitionStyleDisplay struct {
	// The display type for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/dashboard#type Dashboard#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

