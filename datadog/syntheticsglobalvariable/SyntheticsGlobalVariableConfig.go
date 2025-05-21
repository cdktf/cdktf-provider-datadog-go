// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsglobalvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsGlobalVariableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Synthetics global variable name. Must be all uppercase with underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#name SyntheticsGlobalVariable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the global variable. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#description SyntheticsGlobalVariable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set to true, the global variable is a FIDO variable. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#is_fido SyntheticsGlobalVariable#is_fido}
	IsFido interface{} `field:"optional" json:"isFido" yaml:"isFido"`
	// If set to true, the global variable is a TOTP variable. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#is_totp SyntheticsGlobalVariable#is_totp}
	IsTotp interface{} `field:"optional" json:"isTotp" yaml:"isTotp"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#options SyntheticsGlobalVariable#options}
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Id of the Synthetics test to use for a variable from test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#parse_test_id SyntheticsGlobalVariable#parse_test_id}
	ParseTestId *string `field:"optional" json:"parseTestId" yaml:"parseTestId"`
	// parse_test_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#parse_test_options SyntheticsGlobalVariable#parse_test_options}
	ParseTestOptions interface{} `field:"optional" json:"parseTestOptions" yaml:"parseTestOptions"`
	// A list of role identifiers to associate with the Synthetics global variable.
	//
	// **Deprecated.** This field is no longer supported by the Datadog API. Please use `datadog_restriction_policy` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#restricted_roles SyntheticsGlobalVariable#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// If set to true, the value of the global variable is hidden.
	//
	// This setting is automatically set to `true` if `is_totp` or `is_fido` is set to `true`. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#secure SyntheticsGlobalVariable#secure}
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
	// A list of tags to associate with your synthetics global variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#tags SyntheticsGlobalVariable#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// The value of the global variable. Required unless `is_fido` is set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/synthetics_global_variable#value SyntheticsGlobalVariable#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

