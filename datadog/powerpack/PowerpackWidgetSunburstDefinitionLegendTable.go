// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetSunburstDefinitionLegendTable struct {
	// The type of legend (table or none). Valid values are `table`, `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#type Powerpack#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

