// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetPowerpackDefinitionTemplateVariables struct {
	// controlled_by_powerpack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#controlled_by_powerpack Dashboard#controlled_by_powerpack}
	ControlledByPowerpack interface{} `field:"optional" json:"controlledByPowerpack" yaml:"controlledByPowerpack"`
	// controlled_externally block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#controlled_externally Dashboard#controlled_externally}
	ControlledExternally interface{} `field:"optional" json:"controlledExternally" yaml:"controlledExternally"`
}

