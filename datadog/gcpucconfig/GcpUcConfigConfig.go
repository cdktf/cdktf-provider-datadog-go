// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gcpucconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GcpUcConfigConfig struct {
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
	// The Google Cloud account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/gcp_uc_config#billing_account_id GcpUcConfig#billing_account_id}
	BillingAccountId *string `field:"required" json:"billingAccountId" yaml:"billingAccountId"`
	// The Google Cloud bucket name used to store the Usage Cost export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/gcp_uc_config#bucket_name GcpUcConfig#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The export dataset name used for the Google Cloud Usage Cost report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/gcp_uc_config#export_dataset_name GcpUcConfig#export_dataset_name}
	ExportDatasetName *string `field:"required" json:"exportDatasetName" yaml:"exportDatasetName"`
	// The name of the Google Cloud Usage Cost report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/gcp_uc_config#export_project_name GcpUcConfig#export_project_name}
	ExportProjectName *string `field:"required" json:"exportProjectName" yaml:"exportProjectName"`
	// The unique Google Cloud service account email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/gcp_uc_config#service_account GcpUcConfig#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// The export prefix used for the Google Cloud Usage Cost report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/gcp_uc_config#export_prefix GcpUcConfig#export_prefix}
	ExportPrefix *string `field:"optional" json:"exportPrefix" yaml:"exportPrefix"`
}

