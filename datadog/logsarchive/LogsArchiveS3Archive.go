// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsarchive


type LogsArchiveS3Archive struct {
	// Your AWS account id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/logs_archive#account_id LogsArchive#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Name of your s3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/logs_archive#bucket LogsArchive#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Your AWS role name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/logs_archive#role_name LogsArchive#role_name}
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
	// Path where the archive is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/logs_archive#path LogsArchive#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

