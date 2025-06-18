// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryTableDefinitionRequestFormulaCellDisplayModeOptions struct {
	// The type of trend line to display. Valid values are `area`, `line`, `bars`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#trend_type Powerpack#trend_type}
	TrendType *string `field:"optional" json:"trendType" yaml:"trendType"`
	// The scale of the y-axis. Valid values are `shared`, `independent`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/powerpack#y_scale Powerpack#y_scale}
	YScale *string `field:"optional" json:"yScale" yaml:"yScale"`
}

