// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionHttpTokenAuthUrlParameter struct {
	// URL parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/action_connection#name ActionConnection#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// URL parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/action_connection#value ActionConnection#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

