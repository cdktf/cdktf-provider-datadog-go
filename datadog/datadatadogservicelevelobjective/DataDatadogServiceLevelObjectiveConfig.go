// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogservicelevelobjective

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogServiceLevelObjectiveConfig struct {
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
	// A SLO ID to limit the search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/data-sources/service_level_objective#id DataDatadogServiceLevelObjective#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Filter results based on SLO numerator and denominator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/data-sources/service_level_objective#metrics_query DataDatadogServiceLevelObjective#metrics_query}
	MetricsQuery *string `field:"optional" json:"metricsQuery" yaml:"metricsQuery"`
	// Filter results based on SLO names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/data-sources/service_level_objective#name_query DataDatadogServiceLevelObjective#name_query}
	NameQuery *string `field:"optional" json:"nameQuery" yaml:"nameQuery"`
	// Filter results based on a single SLO tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.35.0/docs/data-sources/service_level_objective#tags_query DataDatadogServiceLevelObjective#tags_query}
	TagsQuery *string `field:"optional" json:"tagsQuery" yaml:"tagsQuery"`
}

