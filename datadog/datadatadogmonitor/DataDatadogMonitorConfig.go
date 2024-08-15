// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.0/docs/data-sources/monitor#id DataDatadogMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of monitor tags to limit the search. This filters on the tags set on the monitor itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.0/docs/data-sources/monitor#monitor_tags_filter DataDatadogMonitor#monitor_tags_filter}
	MonitorTagsFilter *[]*string `field:"optional" json:"monitorTagsFilter" yaml:"monitorTagsFilter"`
	// A monitor name to limit the search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.0/docs/data-sources/monitor#name_filter DataDatadogMonitor#name_filter}
	NameFilter *string `field:"optional" json:"nameFilter" yaml:"nameFilter"`
	// A list of tags to limit the search. This filters on the monitor scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.43.0/docs/data-sources/monitor#tags_filter DataDatadogMonitor#tags_filter}
	TagsFilter *[]*string `field:"optional" json:"tagsFilter" yaml:"tagsFilter"`
}

