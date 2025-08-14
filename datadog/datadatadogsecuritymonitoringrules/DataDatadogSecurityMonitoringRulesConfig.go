// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogsecuritymonitoringrules

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogSecurityMonitoringRulesConfig struct {
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
	// Limit the search to default rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/data-sources/security_monitoring_rules#default_only_filter DataDatadogSecurityMonitoringRules#default_only_filter}
	DefaultOnlyFilter interface{} `field:"optional" json:"defaultOnlyFilter" yaml:"defaultOnlyFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/data-sources/security_monitoring_rules#id DataDatadogSecurityMonitoringRules#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A rule name to limit the search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/data-sources/security_monitoring_rules#name_filter DataDatadogSecurityMonitoringRules#name_filter}
	NameFilter *string `field:"optional" json:"nameFilter" yaml:"nameFilter"`
	// A list of tags to limit the search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/data-sources/security_monitoring_rules#tags_filter DataDatadogSecurityMonitoringRules#tags_filter}
	TagsFilter *[]*string `field:"optional" json:"tagsFilter" yaml:"tagsFilter"`
	// Limit the search to user rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/data-sources/security_monitoring_rules#user_only_filter DataDatadogSecurityMonitoringRules#user_only_filter}
	UserOnlyFilter interface{} `field:"optional" json:"userOnlyFilter" yaml:"userOnlyFilter"`
}

