// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountLogsConfigLambdaForwarder struct {
	// List of Datadog Lambda Log Forwarder ARNs in your AWS account. Defaults to `[]`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/integration_aws_account#lambdas IntegrationAwsAccount#lambdas}
	Lambdas *[]*string `field:"optional" json:"lambdas" yaml:"lambdas"`
	// List of service IDs set to enable automatic log collection.
	//
	// Use [`datadog_integration_aws_available_logs_services` data source](https://registry.terraform.io/providers/DataDog/datadog/latest/docs/data-sources/integration_aws_available_logs_services) or [the AWS Logs Integration API](https://docs.datadoghq.com/api/latest/aws-logs-integration/?#get-list-of-aws-log-ready-services) to get allowed values. Defaults to `[]`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/integration_aws_account#sources IntegrationAwsAccount#sources}
	Sources *[]*string `field:"optional" json:"sources" yaml:"sources"`
}

