// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogroleusers

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogRoleUsersConfig struct {
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
	// The role's identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/data-sources/role_users#role_id DataDatadogRoleUsers#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// When true, `filter_keyword` string is exact matched against the user's `name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/data-sources/role_users#exact_match DataDatadogRoleUsers#exact_match}
	ExactMatch interface{} `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Search query, can be user name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/data-sources/role_users#filter DataDatadogRoleUsers#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

