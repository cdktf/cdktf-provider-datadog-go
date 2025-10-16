// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package awscurconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AwsCurConfigConfig struct {
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
	// The AWS account ID of your billing/payer account. For AWS Organizations, this is typically the management account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/aws_cur_config#account_id AwsCurConfig#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The S3 bucket name where your AWS Cost and Usage Report files are stored.
	//
	// This bucket must have appropriate IAM permissions for Datadog access and should be in the same AWS account as the CUR report.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/aws_cur_config#bucket_name AwsCurConfig#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The exact name of your AWS Cost and Usage Report as configured in AWS Billing preferences.
	//
	// This must match the report name exactly as it appears in your AWS billing settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/aws_cur_config#report_name AwsCurConfig#report_name}
	ReportName *string `field:"required" json:"reportName" yaml:"reportName"`
	// The S3 key prefix where your Cost and Usage Report files are stored within the bucket (e.g., 'cur-reports/', 'billing/cur/').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/aws_cur_config#report_prefix AwsCurConfig#report_prefix}
	ReportPrefix *string `field:"required" json:"reportPrefix" yaml:"reportPrefix"`
	// account_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/aws_cur_config#account_filters AwsCurConfig#account_filters}
	AccountFilters *AwsCurConfigAccountFilters `field:"optional" json:"accountFilters" yaml:"accountFilters"`
	// The AWS region where the S3 bucket containing your Cost and Usage Report is located (e.g., us-east-1, eu-west-1).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/aws_cur_config#bucket_region AwsCurConfig#bucket_region}
	BucketRegion *string `field:"optional" json:"bucketRegion" yaml:"bucketRegion"`
}

