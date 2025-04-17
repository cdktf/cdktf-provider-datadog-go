// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SensitiveDataScannerGroupConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/sensitive_data_scanner_group#filter SensitiveDataScannerGroup#filter}
	Filter *SensitiveDataScannerGroupFilter `field:"required" json:"filter" yaml:"filter"`
	// Whether or not the scanning group is enabled.
	//
	// If the group doesn't contain any rule or if all the rules in it are disabled, the group is force-disabled by our backend
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/sensitive_data_scanner_group#is_enabled SensitiveDataScannerGroup#is_enabled}
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Name of the Datadog scanning group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/sensitive_data_scanner_group#name SensitiveDataScannerGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// List of products the scanning group applies. Valid values are `logs`, `rum`, `events`, `apm`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/sensitive_data_scanner_group#product_list SensitiveDataScannerGroup#product_list}
	ProductList *[]*string `field:"required" json:"productList" yaml:"productList"`
	// Description of the Datadog scanning group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/sensitive_data_scanner_group#description SensitiveDataScannerGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.0/docs/resources/sensitive_data_scanner_group#id SensitiveDataScannerGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

