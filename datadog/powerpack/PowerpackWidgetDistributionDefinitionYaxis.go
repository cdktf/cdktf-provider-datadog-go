// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetDistributionDefinitionYaxis struct {
	// Always include zero or fit the axis to the data range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#include_zero Powerpack#include_zero}
	IncludeZero interface{} `field:"optional" json:"includeZero" yaml:"includeZero"`
	// The label of the axis to display on the graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#label Powerpack#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Specify the maximum value to show on the Y-axis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#max Powerpack#max}
	Max *string `field:"optional" json:"max" yaml:"max"`
	// Specify the minimum value to show on the Y-axis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#min Powerpack#min}
	Min *string `field:"optional" json:"min" yaml:"min"`
	// Specify the scale type, options: `linear`, `log`, `pow`, `sqrt`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/powerpack#scale Powerpack#scale}
	Scale *string `field:"optional" json:"scale" yaml:"scale"`
}

