// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package userrole

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserRoleConfig struct {
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
	// ID of the role that the user is assigned to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/user_role#role_id UserRole#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// The ID of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/user_role#user_id UserRole#user_id}
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

