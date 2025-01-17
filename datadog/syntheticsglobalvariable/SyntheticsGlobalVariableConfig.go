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
	// Synthetics global variable name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#name SyntheticsGlobalVariable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the global variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#value SyntheticsGlobalVariable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Description of the global variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#description SyntheticsGlobalVariable#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#id SyntheticsGlobalVariable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#options SyntheticsGlobalVariable#options}
	Options *SyntheticsGlobalVariableOptions `field:"optional" json:"options" yaml:"options"`
	// Id of the Synthetics test to use for a variable from test.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#parse_test_id SyntheticsGlobalVariable#parse_test_id}
	ParseTestId *string `field:"optional" json:"parseTestId" yaml:"parseTestId"`
	// parse_test_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#parse_test_options SyntheticsGlobalVariable#parse_test_options}
	ParseTestOptions *SyntheticsGlobalVariableParseTestOptions `field:"optional" json:"parseTestOptions" yaml:"parseTestOptions"`
	// A list of role identifiers to associate with the Synthetics global variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#restricted_roles SyntheticsGlobalVariable#restricted_roles}
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// If set to true, the value of the global variable is hidden. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#secure SyntheticsGlobalVariable#secure}
	Secure interface{} `field:"optional" json:"secure" yaml:"secure"`
	// A list of tags to associate with your synthetics global variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/synthetics_global_variable#tags SyntheticsGlobalVariable#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

