// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountLogsConfig struct {
	// lambda_forwarder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/integration_aws_account#lambda_forwarder IntegrationAwsAccount#lambda_forwarder}
	LambdaForwarder *IntegrationAwsAccountLogsConfigLambdaForwarder `field:"optional" json:"lambdaForwarder" yaml:"lambdaForwarder"`
}

