// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackTemplateVariables struct {
	// The name of the powerpack template variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or many default values for powerpack template variables on load.
	//
	// If more than one default is specified, they will be unioned together with `OR`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/powerpack#defaults Powerpack#defaults}
	Defaults *[]*string `field:"optional" json:"defaults" yaml:"defaults"`
}

