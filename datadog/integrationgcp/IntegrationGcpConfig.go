// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationgcp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationGcpConfig struct {
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
	// Your email found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#client_email IntegrationGcp#client_email}
	ClientEmail *string `field:"required" json:"clientEmail" yaml:"clientEmail"`
	// Your ID found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#client_id IntegrationGcp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Your private key name found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#private_key IntegrationGcp#private_key}
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Your private key ID found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#private_key_id IntegrationGcp#private_key_id}
	PrivateKeyId *string `field:"required" json:"privateKeyId" yaml:"privateKeyId"`
	// Your Google Cloud project ID found in your JSON service account key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#project_id IntegrationGcp#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Silence monitors for expected GCE instance shutdowns. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#automute IntegrationGcp#automute}
	Automute interface{} `field:"optional" json:"automute" yaml:"automute"`
	// Tags to filter which Cloud Run revisions are imported into Datadog. Only revisions that meet specified criteria are monitored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#cloud_run_revision_filters IntegrationGcp#cloud_run_revision_filters}
	CloudRunRevisionFilters *[]*string `field:"optional" json:"cloudRunRevisionFilters" yaml:"cloudRunRevisionFilters"`
	// Whether Datadog collects cloud security posture management resources from your GCP project.
	//
	// If enabled, requires `resource_collection_enabled` to also be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#cspm_resource_collection_enabled IntegrationGcp#cspm_resource_collection_enabled}
	CspmResourceCollectionEnabled interface{} `field:"optional" json:"cspmResourceCollectionEnabled" yaml:"cspmResourceCollectionEnabled"`
	// Limit the GCE instances that are pulled into Datadog by using tags.
	//
	// Only hosts that match one of the defined tags are imported into Datadog. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#host_filters IntegrationGcp#host_filters}
	HostFilters *string `field:"optional" json:"hostFilters" yaml:"hostFilters"`
	// When enabled, Datadog scans for all resource change data in your Google Cloud environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#is_resource_change_collection_enabled IntegrationGcp#is_resource_change_collection_enabled}
	IsResourceChangeCollectionEnabled interface{} `field:"optional" json:"isResourceChangeCollectionEnabled" yaml:"isResourceChangeCollectionEnabled"`
	// When enabled, Datadog will attempt to collect Security Command Center Findings.
	//
	// Note: This requires additional permissions on the service account. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#is_security_command_center_enabled IntegrationGcp#is_security_command_center_enabled}
	IsSecurityCommandCenterEnabled interface{} `field:"optional" json:"isSecurityCommandCenterEnabled" yaml:"isSecurityCommandCenterEnabled"`
	// When enabled, Datadog scans for all resources in your GCP environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_gcp#resource_collection_enabled IntegrationGcp#resource_collection_enabled}
	ResourceCollectionEnabled interface{} `field:"optional" json:"resourceCollectionEnabled" yaml:"resourceCollectionEnabled"`
}

