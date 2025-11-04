// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionRequestStyle struct {
	// The type of lines displayed. Valid values are `dashed`, `dotted`, `solid`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#line_type Powerpack#line_type}
	LineType *string `field:"optional" json:"lineType" yaml:"lineType"`
	// The width of line displayed. Valid values are `normal`, `thick`, `thin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#line_width Powerpack#line_width}
	LineWidth *string `field:"optional" json:"lineWidth" yaml:"lineWidth"`
	// A color palette to apply to the widget. The available options are available at: https://docs.datadoghq.com/dashboards/widgets/timeseries/#appearance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/powerpack#palette Powerpack#palette}
	Palette *string `field:"optional" json:"palette" yaml:"palette"`
}

