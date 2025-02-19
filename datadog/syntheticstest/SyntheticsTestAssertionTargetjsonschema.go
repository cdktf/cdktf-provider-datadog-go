// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticstest


type SyntheticsTestAssertionTargetjsonschema struct {
	// The JSON Schema to validate the body against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/synthetics_test#jsonschema SyntheticsTest#jsonschema}
	Jsonschema *string `field:"required" json:"jsonschema" yaml:"jsonschema"`
	// The meta schema to use for the JSON Schema. Defaults to `"draft-07"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/synthetics_test#metaschema SyntheticsTest#metaschema}
	Metaschema *string `field:"optional" json:"metaschema" yaml:"metaschema"`
}

