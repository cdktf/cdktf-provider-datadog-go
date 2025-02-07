// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountAuthConfig struct {
	// aws_auth_config_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/integration_aws_account#aws_auth_config_keys IntegrationAwsAccount#aws_auth_config_keys}
	AwsAuthConfigKeys *IntegrationAwsAccountAuthConfigAwsAuthConfigKeys `field:"optional" json:"awsAuthConfigKeys" yaml:"awsAuthConfigKeys"`
	// aws_auth_config_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/integration_aws_account#aws_auth_config_role IntegrationAwsAccount#aws_auth_config_role}
	AwsAuthConfigRole *IntegrationAwsAccountAuthConfigAwsAuthConfigRole `field:"optional" json:"awsAuthConfigRole" yaml:"awsAuthConfigRole"`
}

