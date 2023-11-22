// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogteammemberships

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogTeamMembershipsConfig struct {
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
	// The team's identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/data-sources/team_memberships#team_id DataDatadogTeamMemberships#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// When true, `filter_keyword` string is exact matched against the user's `email`, followed by `name`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/data-sources/team_memberships#exact_match DataDatadogTeamMemberships#exact_match}
	ExactMatch interface{} `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Search query, can be user email or name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.33.0/docs/data-sources/team_memberships#filter_keyword DataDatadogTeamMemberships#filter_keyword}
	FilterKeyword *string `field:"optional" json:"filterKeyword" yaml:"filterKeyword"`
}

