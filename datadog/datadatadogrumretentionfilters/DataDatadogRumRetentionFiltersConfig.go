// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogrumretentionfilters

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogRumRetentionFiltersConfig struct {
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
	// RUM application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/data-sources/rum_retention_filters#application_id DataDatadogRumRetentionFilters#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
}

