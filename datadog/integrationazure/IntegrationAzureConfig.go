// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationazure

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationAzureConfig struct {
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
	// Your Azure web application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#client_id IntegrationAzure#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// (Required for Initial Creation) Your Azure web application secret key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#client_secret IntegrationAzure#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// Your Azure Active Directory ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#tenant_name IntegrationAzure#tenant_name}
	TenantName *string `field:"required" json:"tenantName" yaml:"tenantName"`
	// This comma-separated list of tags (in the form `key:value,key:value`) defines a filter that Datadog uses when collecting metrics from Azure App Service Plans.
	//
	// Only App Service Plans that match one of the defined tags are imported into Datadog. The rest, including the apps and functions running on them, are ignored. This also filters the metrics for any App or Function running on the App Service Plan(s). Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#app_service_plan_filters IntegrationAzure#app_service_plan_filters}
	AppServicePlanFilters *string `field:"optional" json:"appServicePlanFilters" yaml:"appServicePlanFilters"`
	// Silence monitors for expected Azure VM shutdowns. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#automute IntegrationAzure#automute}
	Automute interface{} `field:"optional" json:"automute" yaml:"automute"`
	// This comma-separated list of tags (in the form `key:value,key:value`) defines a filter that Datadog uses when collecting metrics from Azure Container Apps.
	//
	// Only Container Apps that match one of the defined tags are imported into Datadog. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#container_app_filters IntegrationAzure#container_app_filters}
	ContainerAppFilters *string `field:"optional" json:"containerAppFilters" yaml:"containerAppFilters"`
	// When enabled, Datadog’s Cloud Security Management product scans resource configurations monitored by this app registration.
	//
	// Note: This requires `resource_collection_enabled` to be set to true. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#cspm_enabled IntegrationAzure#cspm_enabled}
	CspmEnabled interface{} `field:"optional" json:"cspmEnabled" yaml:"cspmEnabled"`
	// Enable custom metrics for your organization. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#custom_metrics_enabled IntegrationAzure#custom_metrics_enabled}
	CustomMetricsEnabled interface{} `field:"optional" json:"customMetricsEnabled" yaml:"customMetricsEnabled"`
	// String of host tag(s) (in the form `key:value,key:value`) defines a filter that Datadog will use when collecting metrics from Azure.
	//
	// Limit the Azure instances that are pulled into Datadog by using tags. Only hosts that match one of the defined tags are imported into Datadog. e.x. `env:production,deploymentgroup:red` Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#host_filters IntegrationAzure#host_filters}
	HostFilters *string `field:"optional" json:"hostFilters" yaml:"hostFilters"`
	// When enabled, Datadog collects metadata and configuration info from cloud resources (such as compute instances, databases, and load balancers) monitored by this app registration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.36.1/docs/resources/integration_azure#resource_collection_enabled IntegrationAzure#resource_collection_enabled}
	ResourceCollectionEnabled interface{} `field:"optional" json:"resourceCollectionEnabled" yaml:"resourceCollectionEnabled"`
}

