// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package user

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserConfig struct {
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
	// Email address for user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/user#email User#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// Whether the user is disabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/user#disabled User#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/user#id User#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User's name. Should be set only for password authentication, as it is overridden by Google or SAML authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/user#name User#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of role IDs to assign to the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/user#roles User#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Whether an invitation email should be sent when the user is created. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.75.0/docs/resources/user#send_user_invitation User#send_user_invitation}
	SendUserInvitation interface{} `field:"optional" json:"sendUserInvitation" yaml:"sendUserInvitation"`
}

