// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teampermissionsetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamPermissionSettingConfig struct {
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
	// The identifier for the action. Valid values are `manage_membership`, `edit`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/team_permission_setting#action TeamPermissionSetting#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// ID of the team the team permission setting is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/team_permission_setting#team_id TeamPermissionSetting#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The action value.
	//
	// Valid values are dependent on the action. `manage_membership` action allows `admins`, `members`, `organization`, `user_access_manage` values. `edit` action allows `admins`, `members`, `teams_manage` values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/team_permission_setting#value TeamPermissionSetting#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

