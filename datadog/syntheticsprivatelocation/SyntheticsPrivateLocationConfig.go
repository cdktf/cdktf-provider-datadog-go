// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsprivatelocation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SyntheticsPrivateLocationConfig struct {
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
	// Synthetics private location name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/synthetics_private_location#name SyntheticsPrivateLocation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the private location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/synthetics_private_location#description SyntheticsPrivateLocation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/synthetics_private_location#id SyntheticsPrivateLocation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/synthetics_private_location#metadata SyntheticsPrivateLocation#metadata}
	Metadata *SyntheticsPrivateLocationMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// A list of tags to associate with your synthetics private location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.46.0/docs/resources/synthetics_private_location#tags SyntheticsPrivateLocation#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

