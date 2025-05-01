// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountTracesConfig struct {
	// xray_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.61.0/docs/resources/integration_aws_account#xray_services IntegrationAwsAccount#xray_services}
	XrayServices *IntegrationAwsAccountTracesConfigXrayServices `field:"optional" json:"xrayServices" yaml:"xrayServices"`
}

