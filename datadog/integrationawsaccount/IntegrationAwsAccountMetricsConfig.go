// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount


type IntegrationAwsAccountMetricsConfig struct {
	// Enable EC2 automute for AWS metrics Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/integration_aws_account#automute_enabled IntegrationAwsAccount#automute_enabled}
	AutomuteEnabled interface{} `field:"optional" json:"automuteEnabled" yaml:"automuteEnabled"`
	// Enable CloudWatch alarms collection Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/integration_aws_account#collect_cloudwatch_alarms IntegrationAwsAccount#collect_cloudwatch_alarms}
	CollectCloudwatchAlarms interface{} `field:"optional" json:"collectCloudwatchAlarms" yaml:"collectCloudwatchAlarms"`
	// Enable custom metrics collection Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/integration_aws_account#collect_custom_metrics IntegrationAwsAccount#collect_custom_metrics}
	CollectCustomMetrics interface{} `field:"optional" json:"collectCustomMetrics" yaml:"collectCustomMetrics"`
	// Enable AWS metrics collection Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/integration_aws_account#enabled IntegrationAwsAccount#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// namespace_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/integration_aws_account#namespace_filters IntegrationAwsAccount#namespace_filters}
	NamespaceFilters *IntegrationAwsAccountMetricsConfigNamespaceFilters `field:"optional" json:"namespaceFilters" yaml:"namespaceFilters"`
	// tag_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/integration_aws_account#tag_filters IntegrationAwsAccount#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

