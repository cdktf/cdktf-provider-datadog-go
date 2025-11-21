// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationgcpsts

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationGcpStsConfig struct {
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
	// Your service account email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#client_email IntegrationGcpSts#client_email}
	ClientEmail *string `field:"required" json:"clientEmail" yaml:"clientEmail"`
	// Tags to be associated with GCP metrics and service checks from your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#account_tags IntegrationGcpSts#account_tags}
	AccountTags *[]*string `field:"optional" json:"accountTags" yaml:"accountTags"`
	// Silence monitors for expected GCE instance shutdowns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#automute IntegrationGcpSts#automute}
	Automute interface{} `field:"optional" json:"automute" yaml:"automute"`
	// List of filters to limit the Cloud Run revisions that are pulled into Datadog by using tags.
	//
	// Only Cloud Run revision resources that apply to specified filters are imported into Datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#cloud_run_revision_filters IntegrationGcpSts#cloud_run_revision_filters}
	CloudRunRevisionFilters *[]*string `field:"optional" json:"cloudRunRevisionFilters" yaml:"cloudRunRevisionFilters"`
	// List of filters to limit the VM instances that are pulled into Datadog by using tags.
	//
	// Only VM instance resources that apply to specified filters are imported into Datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#host_filters IntegrationGcpSts#host_filters}
	HostFilters *[]*string `field:"optional" json:"hostFilters" yaml:"hostFilters"`
	// Whether Datadog collects cloud security posture management resources from your GCP project.
	//
	// If enabled, requires `resource_collection_enabled` to also be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#is_cspm_enabled IntegrationGcpSts#is_cspm_enabled}
	IsCspmEnabled interface{} `field:"optional" json:"isCspmEnabled" yaml:"isCspmEnabled"`
	// When enabled, Datadog includes the `X-Goog-User-Project` header to attribute Google Cloud billing and quota usage to the monitored project instead of the default service account project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#is_per_project_quota_enabled IntegrationGcpSts#is_per_project_quota_enabled}
	IsPerProjectQuotaEnabled interface{} `field:"optional" json:"isPerProjectQuotaEnabled" yaml:"isPerProjectQuotaEnabled"`
	// When enabled, Datadog scans for all resource change data in your Google Cloud environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#is_resource_change_collection_enabled IntegrationGcpSts#is_resource_change_collection_enabled}
	IsResourceChangeCollectionEnabled interface{} `field:"optional" json:"isResourceChangeCollectionEnabled" yaml:"isResourceChangeCollectionEnabled"`
	// When enabled, Datadog will attempt to collect Security Command Center Findings.
	//
	// Note: This requires additional permissions on the service account. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#is_security_command_center_enabled IntegrationGcpSts#is_security_command_center_enabled}
	IsSecurityCommandCenterEnabled interface{} `field:"optional" json:"isSecurityCommandCenterEnabled" yaml:"isSecurityCommandCenterEnabled"`
	// Configurations for GCP metric namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#metric_namespace_configs IntegrationGcpSts#metric_namespace_configs}
	MetricNamespaceConfigs interface{} `field:"optional" json:"metricNamespaceConfigs" yaml:"metricNamespaceConfigs"`
	// Configurations for GCP monitored resources. Only monitored resources that apply to specified filters are imported into Datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#monitored_resource_configs IntegrationGcpSts#monitored_resource_configs}
	MonitoredResourceConfigs interface{} `field:"optional" json:"monitoredResourceConfigs" yaml:"monitoredResourceConfigs"`
	// When enabled, Datadog scans for all resources in your GCP environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.81.0/docs/resources/integration_gcp_sts#resource_collection_enabled IntegrationGcpSts#resource_collection_enabled}
	ResourceCollectionEnabled interface{} `field:"optional" json:"resourceCollectionEnabled" yaml:"resourceCollectionEnabled"`
}

