// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetSplitGraphDefinitionSourceWidgetDefinitionTimeseriesDefinitionCustomLink struct {
	// The flag for toggling context menu link visibility.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/dashboard#is_hidden Dashboard#is_hidden}
	IsHidden interface{} `field:"optional" json:"isHidden" yaml:"isHidden"`
	// The label for the custom link URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/dashboard#label Dashboard#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// The URL of the custom link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/dashboard#link Dashboard#link}
	Link *string `field:"optional" json:"link" yaml:"link"`
	// The label ID that refers to a context menu link item.
	//
	// When `override_label` is provided, the client request omits the label field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/dashboard#override_label Dashboard#override_label}
	OverrideLabel *string `field:"optional" json:"overrideLabel" yaml:"overrideLabel"`
}

