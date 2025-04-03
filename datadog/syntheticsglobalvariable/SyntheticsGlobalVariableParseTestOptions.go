// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsglobalvariable


type SyntheticsGlobalVariableParseTestOptions struct {
	// Defines the source to use to extract the value. Valid values are `http_body`, `http_header`, `http_status_code`, `local_variable`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/synthetics_global_variable#type SyntheticsGlobalVariable#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Required when type = `http_header`. Defines the header to use to extract the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/synthetics_global_variable#field SyntheticsGlobalVariable#field}
	Field *string `field:"optional" json:"field" yaml:"field"`
	// When type is `local_variable`, name of the local variable to use to extract the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/synthetics_global_variable#local_variable_name SyntheticsGlobalVariable#local_variable_name}
	LocalVariableName *string `field:"optional" json:"localVariableName" yaml:"localVariableName"`
	// parser block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/resources/synthetics_global_variable#parser SyntheticsGlobalVariable#parser}
	Parser interface{} `field:"optional" json:"parser" yaml:"parser"`
}

