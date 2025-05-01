// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package role

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RoleConfig struct {
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
	// Name of the role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/role#name Role#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// If set to `true`, the role does not have default (restricted) permissions unless they are explicitly set.
	//
	// The `include_restricted` attribute for the `datadog_permissions` data source must be set to `true` to manage default permissions in Terraform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/role#default_permissions_opt_out Role#default_permissions_opt_out}
	DefaultPermissionsOptOut interface{} `field:"optional" json:"defaultPermissionsOptOut" yaml:"defaultPermissionsOptOut"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/role#id Role#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/role#permission Role#permission}
	Permission interface{} `field:"optional" json:"permission" yaml:"permission"`
	// If set to `false`, skip the validation call done during plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/role#validate Role#validate}
	Validate interface{} `field:"optional" json:"validate" yaml:"validate"`
}

