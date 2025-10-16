// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package azureucconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AzureUcConfigConfig struct {
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
	// The tenant ID of the Azure account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/azure_uc_config#account_id AzureUcConfig#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// actual_bill_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/azure_uc_config#actual_bill_config AzureUcConfig#actual_bill_config}
	ActualBillConfig *AzureUcConfigActualBillConfig `field:"required" json:"actualBillConfig" yaml:"actualBillConfig"`
	// amortized_bill_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/azure_uc_config#amortized_bill_config AzureUcConfig#amortized_bill_config}
	AmortizedBillConfig *AzureUcConfigAmortizedBillConfig `field:"required" json:"amortizedBillConfig" yaml:"amortizedBillConfig"`
	// The client ID of the Azure account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/azure_uc_config#client_id AzureUcConfig#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The scope of your observed subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.76.0/docs/resources/azure_uc_config#scope AzureUcConfig#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

