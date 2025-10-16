// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardTemplateVariablePresetTemplateVariable struct {
	// The name of the template variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#name Dashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value that should be assumed by the template variable in this preset.
	//
	// Cannot be used in conjunction with `values`. **Deprecated.** Use `values` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#value Dashboard#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// One or many template variable values within the saved view, which will be unioned together using `OR` if more than one is specified.
	//
	// Cannot be used in conjunction with `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/dashboard#values Dashboard#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

