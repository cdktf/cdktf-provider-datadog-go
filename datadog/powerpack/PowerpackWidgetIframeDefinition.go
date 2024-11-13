// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetIframeDefinition struct {
	// The URL to use as a data source for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/powerpack#url Powerpack#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

