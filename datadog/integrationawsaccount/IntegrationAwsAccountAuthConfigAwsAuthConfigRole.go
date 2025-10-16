// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountAuthConfigAwsAuthConfigRole struct {
	// AWS IAM External ID for associated role. If omitted, one will be generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/integration_aws_account#external_id IntegrationAwsAccount#external_id}
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// AWS IAM Role name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/integration_aws_account#role_name IntegrationAwsAccount#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
}

