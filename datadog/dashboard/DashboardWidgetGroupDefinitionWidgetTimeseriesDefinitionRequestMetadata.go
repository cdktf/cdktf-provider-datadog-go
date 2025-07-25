// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboard


type DashboardWidgetGroupDefinitionWidgetTimeseriesDefinitionRequestMetadata struct {
	// The expression name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#expression Dashboard#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/dashboard#alias_name Dashboard#alias_name}
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

