// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetListStreamDefinitionRequestQueryGroupBy struct {
	// Facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/powerpack#facet Powerpack#facet}
	Facet *string `field:"required" json:"facet" yaml:"facet"`
}

