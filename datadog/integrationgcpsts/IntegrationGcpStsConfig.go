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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/integration_gcp_sts#client_email IntegrationGcpSts#client_email}
	ClientEmail *string `field:"required" json:"clientEmail" yaml:"clientEmail"`
	// Silence monitors for expected GCE instance shutdowns.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/integration_gcp_sts#automute IntegrationGcpSts#automute}
	Automute interface{} `field:"optional" json:"automute" yaml:"automute"`
	// Your Host Filters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/integration_gcp_sts#host_filters IntegrationGcpSts#host_filters}
	HostFilters *[]*string `field:"optional" json:"hostFilters" yaml:"hostFilters"`
	// When enabled, Datadog performs configuration checks across your Google Cloud environment by continuously scanning every resource, which may incur additional charges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.37.0/docs/resources/integration_gcp_sts#is_cspm_enabled IntegrationGcpSts#is_cspm_enabled}
	IsCspmEnabled interface{} `field:"optional" json:"isCspmEnabled" yaml:"isCspmEnabled"`
}

