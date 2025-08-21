// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dashboardjson

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardJsonConfig struct {
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
	// The JSON formatted definition of the Dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard_json#dashboard DashboardJson#dashboard}
	Dashboard *string `field:"required" json:"dashboard" yaml:"dashboard"`
	// A list of dashboard lists this dashboard belongs to.
	//
	// This attribute should not be set if managing the corresponding dashboard lists using Terraform as it causes inconsistent behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard_json#dashboard_lists DashboardJson#dashboard_lists}
	DashboardLists *[]*float64 `field:"optional" json:"dashboardLists" yaml:"dashboardLists"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard_json#id DashboardJson#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The URL of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.72.0/docs/resources/dashboard_json#url DashboardJson#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

