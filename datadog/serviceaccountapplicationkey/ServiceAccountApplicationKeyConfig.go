// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package serviceaccountapplicationkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceAccountApplicationKeyConfig struct {
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
	// Name of the application key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/service_account_application_key#name ServiceAccountApplicationKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// ID of the service account that owns this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/service_account_application_key#service_account_id ServiceAccountApplicationKey#service_account_id}
	ServiceAccountId *string `field:"required" json:"serviceAccountId" yaml:"serviceAccountId"`
	// Authorization scopes for the Application Key. Application Keys configured with no scopes have full access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/service_account_application_key#scopes ServiceAccountApplicationKey#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

