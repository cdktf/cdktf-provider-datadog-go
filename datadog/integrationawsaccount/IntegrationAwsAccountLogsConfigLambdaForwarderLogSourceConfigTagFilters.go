// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountLogsConfigLambdaForwarderLogSourceConfigTagFilters struct {
	// The AWS service for which the tag filters defined in `tags` will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/integration_aws_account#source IntegrationAwsAccount#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The AWS resource tags to filter on for the service specified by `source`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/integration_aws_account#tags IntegrationAwsAccount#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
}

