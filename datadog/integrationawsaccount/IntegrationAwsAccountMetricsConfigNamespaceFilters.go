// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountMetricsConfigNamespaceFilters struct {
	// Exclude only these namespaces from metrics collection.
	//
	// Use [`datadog_integration_aws_available_namespaces` data source](https://registry.terraform.io/providers/DataDog/datadog/latest/docs/data-sources/integration_aws_available_namespaces) to get allowed values. Defaults to `["AWS/SQS", "AWS/ElasticMapReduce", "AWS/Usage"]`. `AWS/SQS`, `AWS/ElasticMapReduce`, and `AWS/Usage` are excluded by default to reduce your AWS CloudWatch costs from `GetMetricData` API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/integration_aws_account#exclude_only IntegrationAwsAccount#exclude_only}
	ExcludeOnly *[]*string `field:"optional" json:"excludeOnly" yaml:"excludeOnly"`
	// Include only these namespaces for metrics collection. Use [`datadog_integration_aws_available_namespaces` data source](https://registry.terraform.io/providers/DataDog/datadog/latest/docs/data-sources/integration_aws_available_namespaces) to get allowed values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/integration_aws_account#include_only IntegrationAwsAccount#include_only}
	IncludeOnly *[]*string `field:"optional" json:"includeOnly" yaml:"includeOnly"`
}

