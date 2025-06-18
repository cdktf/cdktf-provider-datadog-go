// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationfastlyservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationFastlyServiceConfig struct {
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
	// The ID of the Fastly service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_fastly_service#service_id IntegrationFastlyService#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Fastly Account id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_fastly_service#account_id IntegrationFastlyService#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// A list of tags for the Fastly service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_fastly_service#tags IntegrationFastlyService#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

