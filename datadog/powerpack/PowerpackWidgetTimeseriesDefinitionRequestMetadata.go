// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionRequestMetadata struct {
	// The expression name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#expression Powerpack#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The expression alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/powerpack#alias_name Powerpack#alias_name}
	AliasName *string `field:"optional" json:"aliasName" yaml:"aliasName"`
}

