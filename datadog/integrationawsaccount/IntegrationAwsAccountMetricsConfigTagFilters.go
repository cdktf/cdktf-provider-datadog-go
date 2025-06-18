// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountMetricsConfigTagFilters struct {
	// The AWS service for which the tag filters defined in `tags` will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_aws_account#namespace IntegrationAwsAccount#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// The AWS resource tags to filter on for the service specified by `namespace`. Defaults to `[]`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_aws_account#tags IntegrationAwsAccount#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

