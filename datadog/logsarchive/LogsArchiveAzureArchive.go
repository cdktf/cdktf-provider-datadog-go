// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsarchive


type LogsArchiveAzureArchive struct {
	// Your client id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/logs_archive#client_id LogsArchive#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The container where the archive is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/logs_archive#container LogsArchive#container}
	Container *string `field:"required" json:"container" yaml:"container"`
	// The associated storage account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/logs_archive#storage_account LogsArchive#storage_account}
	StorageAccount *string `field:"required" json:"storageAccount" yaml:"storageAccount"`
	// Your tenant id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/logs_archive#tenant_id LogsArchive#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The path where the archive is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.38.0/docs/resources/logs_archive#path LogsArchive#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

