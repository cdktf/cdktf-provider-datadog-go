// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apmretentionfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApmRetentionFilterConfig struct {
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
	// the status of the retention filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/apm_retention_filter#enabled ApmRetentionFilter#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The type of the retention filter, currently only spans-processing-sampling is available. Valid values are `spans-sampling-processor`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/apm_retention_filter#filter_type ApmRetentionFilter#filter_type}
	FilterType *string `field:"required" json:"filterType" yaml:"filterType"`
	// The name of the retention filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/apm_retention_filter#name ApmRetentionFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sample rate to apply to spans going through this retention filter as a string, a value of 1.0 keeps all spans matching the query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/apm_retention_filter#rate ApmRetentionFilter#rate}
	Rate *string `field:"required" json:"rate" yaml:"rate"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/apm_retention_filter#filter ApmRetentionFilter#filter}
	Filter *ApmRetentionFilterFilter `field:"optional" json:"filter" yaml:"filter"`
}

