// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionHttpTokenAuthToken struct {
	// Token name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/action_connection#name ActionConnection#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Token type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/action_connection#type ActionConnection#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Token value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/action_connection#value ActionConnection#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

