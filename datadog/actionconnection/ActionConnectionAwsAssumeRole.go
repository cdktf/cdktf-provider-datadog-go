// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionconnection


type ActionConnectionAwsAssumeRole struct {
	// AWS account that the connection is created for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/action_connection#account_id ActionConnection#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Role to assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/action_connection#role ActionConnection#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

