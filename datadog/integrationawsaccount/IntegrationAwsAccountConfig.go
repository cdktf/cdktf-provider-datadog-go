// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationAwsAccountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Your AWS Account ID without dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#aws_account_id IntegrationAwsAccount#aws_account_id}
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// AWS Account partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#aws_partition IntegrationAwsAccount#aws_partition}
	AwsPartition *string `field:"required" json:"awsPartition" yaml:"awsPartition"`
	// Tags to apply to all metrics in the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#account_tags IntegrationAwsAccount#account_tags}
	AccountTags *[]*string `field:"optional" json:"accountTags" yaml:"accountTags"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#auth_config IntegrationAwsAccount#auth_config}
	AuthConfig *IntegrationAwsAccountAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// aws_regions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#aws_regions IntegrationAwsAccount#aws_regions}
	AwsRegions *IntegrationAwsAccountAwsRegions `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// logs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#logs_config IntegrationAwsAccount#logs_config}
	LogsConfig *IntegrationAwsAccountLogsConfig `field:"optional" json:"logsConfig" yaml:"logsConfig"`
	// metrics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#metrics_config IntegrationAwsAccount#metrics_config}
	MetricsConfig *IntegrationAwsAccountMetricsConfig `field:"optional" json:"metricsConfig" yaml:"metricsConfig"`
	// resources_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#resources_config IntegrationAwsAccount#resources_config}
	ResourcesConfig *IntegrationAwsAccountResourcesConfig `field:"optional" json:"resourcesConfig" yaml:"resourcesConfig"`
	// traces_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.57.0/docs/resources/integration_aws_account#traces_config IntegrationAwsAccount#traces_config}
	TracesConfig *IntegrationAwsAccountTracesConfig `field:"optional" json:"tracesConfig" yaml:"tracesConfig"`
}

