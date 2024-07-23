// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadoghosts

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogHostsConfig struct {
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
	// String to filter search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/data-sources/hosts#filter DataDatadogHosts#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Number of seconds since UNIX epoch from which you want to search your hosts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/data-sources/hosts#from DataDatadogHosts#from}
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// Include information on the muted status of hosts and when the mute expires.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/data-sources/hosts#include_muted_hosts_data DataDatadogHosts#include_muted_hosts_data}
	IncludeMutedHostsData interface{} `field:"optional" json:"includeMutedHostsData" yaml:"includeMutedHostsData"`
	// Direction of sort.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/data-sources/hosts#sort_dir DataDatadogHosts#sort_dir}
	SortDir *string `field:"optional" json:"sortDir" yaml:"sortDir"`
	// Sort hosts by this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/data-sources/hosts#sort_field DataDatadogHosts#sort_field}
	SortField *string `field:"optional" json:"sortField" yaml:"sortField"`
}

