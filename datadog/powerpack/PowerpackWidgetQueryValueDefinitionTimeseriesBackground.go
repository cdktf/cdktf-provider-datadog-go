// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryValueDefinitionTimeseriesBackground struct {
	// Whether the Timeseries is made using an area or bars. Valid values are `bars`, `area`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/powerpack#type Powerpack#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// yaxis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/powerpack#yaxis Powerpack#yaxis}
	Yaxis *PowerpackWidgetQueryValueDefinitionTimeseriesBackgroundYaxis `field:"optional" json:"yaxis" yaml:"yaxis"`
}

