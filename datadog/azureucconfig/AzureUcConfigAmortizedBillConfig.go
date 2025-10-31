// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azureucconfig


type AzureUcConfigAmortizedBillConfig struct {
	// The name of the configured Azure Export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/azure_uc_config#export_name AzureUcConfig#export_name}
	ExportName *string `field:"required" json:"exportName" yaml:"exportName"`
	// The path where the Azure Export is saved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/azure_uc_config#export_path AzureUcConfig#export_path}
	ExportPath *string `field:"required" json:"exportPath" yaml:"exportPath"`
	// The name of the storage account where the Azure Export is saved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/azure_uc_config#storage_account AzureUcConfig#storage_account}
	StorageAccount *string `field:"required" json:"storageAccount" yaml:"storageAccount"`
	// The name of the storage container where the Azure Export is saved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.78.0/docs/resources/azure_uc_config#storage_container AzureUcConfig#storage_container}
	StorageContainer *string `field:"required" json:"storageContainer" yaml:"storageContainer"`
}

