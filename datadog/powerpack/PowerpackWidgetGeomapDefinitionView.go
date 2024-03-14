// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetGeomapDefinitionView struct {
	// The two-letter ISO code of a country to focus the map on (or `WORLD`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/powerpack#focus Powerpack#focus}
	Focus *string `field:"required" json:"focus" yaml:"focus"`
}

