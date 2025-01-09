// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountResourcesConfig struct {
	// Enable Cloud Security Management to scan AWS resources for vulnerabilities, misconfigurations, identity risks, and compliance violations.
	//
	// Requires `extended_collection` to be set to `true`. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/integration_aws_account#cloud_security_posture_management_collection IntegrationAwsAccount#cloud_security_posture_management_collection}
	CloudSecurityPostureManagementCollection interface{} `field:"optional" json:"cloudSecurityPostureManagementCollection" yaml:"cloudSecurityPostureManagementCollection"`
	// Whether Datadog collects additional attributes and configuration information about the resources in your AWS account.
	//
	// Required for `cloud_security_posture_management_collection`. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/integration_aws_account#extended_collection IntegrationAwsAccount#extended_collection}
	ExtendedCollection interface{} `field:"optional" json:"extendedCollection" yaml:"extendedCollection"`
}

