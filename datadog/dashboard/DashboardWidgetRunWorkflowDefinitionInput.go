// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetRunWorkflowDefinitionInput struct {
	// Name of the workflow input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Dashboard template variable. Can be suffixed with `.value` or `.key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard#value Dashboard#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

