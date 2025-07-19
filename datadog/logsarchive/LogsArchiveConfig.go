// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsarchive

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsArchiveConfig struct {
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
	// Your archive name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#name LogsArchive#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The archive query/filter. Logs matching this query are included in the archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#query LogsArchive#query}
	Query *string `field:"required" json:"query" yaml:"query"`
	// azure_archive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#azure_archive LogsArchive#azure_archive}
	AzureArchive *LogsArchiveAzureArchive `field:"optional" json:"azureArchive" yaml:"azureArchive"`
	// gcs_archive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#gcs_archive LogsArchive#gcs_archive}
	GcsArchive *LogsArchiveGcsArchive `field:"optional" json:"gcsArchive" yaml:"gcsArchive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#id LogsArchive#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// To store the tags in the archive, set the value `true`.
	//
	// If it is set to `false`, the tags will be dropped when the logs are sent to the archive. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#include_tags LogsArchive#include_tags}
	IncludeTags interface{} `field:"optional" json:"includeTags" yaml:"includeTags"`
	// To limit the rehydration scan size for the archive, set a value in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#rehydration_max_scan_size_in_gb LogsArchive#rehydration_max_scan_size_in_gb}
	RehydrationMaxScanSizeInGb *float64 `field:"optional" json:"rehydrationMaxScanSizeInGb" yaml:"rehydrationMaxScanSizeInGb"`
	// An array of tags to add to rehydrated logs from an archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#rehydration_tags LogsArchive#rehydration_tags}
	RehydrationTags *[]*string `field:"optional" json:"rehydrationTags" yaml:"rehydrationTags"`
	// s3_archive block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.68.0/docs/resources/logs_archive#s3_archive LogsArchive#s3_archive}
	S3Archive *LogsArchiveS3Archive `field:"optional" json:"s3Archive" yaml:"s3Archive"`
}

