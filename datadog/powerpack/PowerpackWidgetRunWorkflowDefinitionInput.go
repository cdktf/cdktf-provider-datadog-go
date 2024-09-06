// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetRunWorkflowDefinitionInput struct {
	// Name of the workflow input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dashboard template variable. Can be suffixed with `.value` or `.key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.0/docs/resources/powerpack#value Powerpack#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

