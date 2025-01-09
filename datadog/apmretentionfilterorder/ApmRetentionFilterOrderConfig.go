// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apmretentionfilterorder

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApmRetentionFilterOrderConfig struct {
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
	// The filter IDs list. The order of filters IDs in this attribute defines the overall APM retention filters order.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.51.0/docs/resources/apm_retention_filter_order#filter_ids ApmRetentionFilterOrder#filter_ids}
	FilterIds *[]*string `field:"required" json:"filterIds" yaml:"filterIds"`
}

