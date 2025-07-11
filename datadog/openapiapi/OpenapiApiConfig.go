// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package openapiapi

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OpenapiApiConfig struct {
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
	// The textual content of the OpenAPI specification.
	//
	// Use [`file()`](https://developer.hashicorp.com/terraform/language/functions/file) in order to reference another file in the repository (see exmaple).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/openapi_api#spec OpenapiApi#spec}
	Spec *string `field:"required" json:"spec" yaml:"spec"`
}

