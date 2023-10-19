// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogserviceaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogServiceAccountConfig struct {
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
	// Filter all users and service accounts by name, email, or role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.31.0/docs/data-sources/service_account#filter DataDatadogServiceAccount#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Filter on status attribute. Comma separated list, with possible values `Active`, `Pending`, and `Disabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.31.0/docs/data-sources/service_account#filter_status DataDatadogServiceAccount#filter_status}
	FilterStatus *string `field:"optional" json:"filterStatus" yaml:"filterStatus"`
	// The service account's ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.31.0/docs/data-sources/service_account#id DataDatadogServiceAccount#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

