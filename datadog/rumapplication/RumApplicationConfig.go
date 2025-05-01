// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rumapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RumApplicationConfig struct {
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
	// Name of the RUM application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/rum_application#name RumApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Type of the RUM application. Supported values are `browser`, `ios`, `android`, `react-native`, `flutter`. Defaults to `"browser"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/rum_application#type RumApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

