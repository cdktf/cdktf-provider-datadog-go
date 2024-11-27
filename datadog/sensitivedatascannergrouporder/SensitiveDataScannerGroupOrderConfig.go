// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sensitivedatascannergrouporder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SensitiveDataScannerGroupOrderConfig struct {
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
	// The list of Sensitive Data Scanner group IDs, in order.
	//
	// Logs are tested against the query filter of each index one by one following the order of the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.49.0/docs/resources/sensitive_data_scanner_group_order#group_ids SensitiveDataScannerGroupOrder#group_ids}
	GroupIds *[]*string `field:"required" json:"groupIds" yaml:"groupIds"`
}

