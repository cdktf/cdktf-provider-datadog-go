// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationaws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationAwsConfig struct {
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
	// Your AWS access key ID. Only required if your AWS account is a GovCloud or China account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#access_key_id IntegrationAws#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// Your AWS Account ID without dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#account_id IntegrationAws#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Enables or disables metric collection for specific AWS namespaces for this AWS account only.
	//
	// A list of namespaces can be found at the [available namespace rules API endpoint](https://docs.datadoghq.com/api/v1/aws-integration/#list-namespace-rules).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#account_specific_namespace_rules IntegrationAws#account_specific_namespace_rules}
	AccountSpecificNamespaceRules *map[string]interface{} `field:"optional" json:"accountSpecificNamespaceRules" yaml:"accountSpecificNamespaceRules"`
	// Whether Datadog collects cloud security posture management resources from your AWS account.
	//
	// This includes additional resources not covered under the general resource_collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#cspm_resource_collection_enabled IntegrationAws#cspm_resource_collection_enabled}
	CspmResourceCollectionEnabled *string `field:"optional" json:"cspmResourceCollectionEnabled" yaml:"cspmResourceCollectionEnabled"`
	// An array of AWS regions to exclude from metrics collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#excluded_regions IntegrationAws#excluded_regions}
	ExcludedRegions *[]*string `field:"optional" json:"excludedRegions" yaml:"excludedRegions"`
	// Whether Datadog collects additional attributes and configuration information about the resources in your AWS account. Required for `cspm_resource_collection_enabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#extended_resource_collection_enabled IntegrationAws#extended_resource_collection_enabled}
	ExtendedResourceCollectionEnabled *string `field:"optional" json:"extendedResourceCollectionEnabled" yaml:"extendedResourceCollectionEnabled"`
	// Array of EC2 tags (in the form `key:value`) defines a filter that Datadog uses when collecting metrics from EC2.
	//
	// Wildcards, such as `?` (for single characters) and `*` (for multiple characters) can also be used. Only hosts that match one of the defined tags will be imported into Datadog. The rest will be ignored. Host matching a given tag can also be excluded by adding `!` before the tag. e.x. `env:production,instance-type:c1.*,!region:us-east-1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#filter_tags IntegrationAws#filter_tags}
	FilterTags *[]*string `field:"optional" json:"filterTags" yaml:"filterTags"`
	// Array of tags (in the form `key:value`) to add to all hosts and metrics reporting through this integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#host_tags IntegrationAws#host_tags}
	HostTags *[]*string `field:"optional" json:"hostTags" yaml:"hostTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#id IntegrationAws#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether Datadog collects metrics for this AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#metrics_collection_enabled IntegrationAws#metrics_collection_enabled}
	MetricsCollectionEnabled *string `field:"optional" json:"metricsCollectionEnabled" yaml:"metricsCollectionEnabled"`
	// Whether Datadog collects a standard set of resources from your AWS account. **Deprecated.** Deprecated in favor of `extended_resource_collection_enabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#resource_collection_enabled IntegrationAws#resource_collection_enabled}
	ResourceCollectionEnabled *string `field:"optional" json:"resourceCollectionEnabled" yaml:"resourceCollectionEnabled"`
	// Your Datadog role delegation name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#role_name IntegrationAws#role_name}
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Your AWS secret access key. Only required if your AWS account is a GovCloud or China account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.48.0/docs/resources/integration_aws#secret_access_key IntegrationAws#secret_access_key}
	SecretAccessKey *string `field:"optional" json:"secretAccessKey" yaml:"secretAccessKey"`
}

