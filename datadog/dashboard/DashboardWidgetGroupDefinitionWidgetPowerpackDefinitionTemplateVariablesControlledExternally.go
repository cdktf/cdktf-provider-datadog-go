// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetPowerpackDefinitionTemplateVariablesControlledExternally struct {
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// One or many template variable values within the saved view, which will be unioned together using `OR` if more than one is specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#values Dashboard#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable dropdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.50.0/docs/resources/dashboard#prefix Dashboard#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

