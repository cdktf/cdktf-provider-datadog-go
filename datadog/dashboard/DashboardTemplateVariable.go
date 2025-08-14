// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardTemplateVariable struct {
	// The name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of values that the template variable drop-down is be limited to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#available_values Dashboard#available_values}
	AvailableValues *[]*string `field:"optional" json:"availableValues" yaml:"availableValues"`
	// The default value for the template variable on dashboard load.
	//
	// Cannot be used in conjunction with `defaults`. **Deprecated.** Use `defaults` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#default Dashboard#default}
	Default *string `field:"optional" json:"default" yaml:"default"`
	// One or many default values for template variables on load.
	//
	// If more than one default is specified, they will be unioned together with `OR`. Cannot be used in conjunction with `default`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#defaults Dashboard#defaults}
	Defaults *[]*string `field:"optional" json:"defaults" yaml:"defaults"`
	// The tag prefix associated with the variable. Only tags with this prefix appear in the variable dropdown.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/dashboard#prefix Dashboard#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

