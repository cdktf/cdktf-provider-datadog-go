// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datadatadogservicelevelobjectives

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataDatadogServiceLevelObjectivesConfig struct {
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
	// Throw an error if no results are found. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/data-sources/service_level_objectives#error_on_empty_result DataDatadogServiceLevelObjectives#error_on_empty_result}
	ErrorOnEmptyResult interface{} `field:"optional" json:"errorOnEmptyResult" yaml:"errorOnEmptyResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/data-sources/service_level_objectives#id DataDatadogServiceLevelObjectives#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An array of SLO IDs to limit the search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/data-sources/service_level_objectives#ids DataDatadogServiceLevelObjectives#ids}
	Ids *[]*string `field:"optional" json:"ids" yaml:"ids"`
	// Filter results based on SLO numerator and denominator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/data-sources/service_level_objectives#metrics_query DataDatadogServiceLevelObjectives#metrics_query}
	MetricsQuery *string `field:"optional" json:"metricsQuery" yaml:"metricsQuery"`
	// Filter results based on SLO names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/data-sources/service_level_objectives#name_query DataDatadogServiceLevelObjectives#name_query}
	NameQuery *string `field:"optional" json:"nameQuery" yaml:"nameQuery"`
	// The query string to filter results based on SLO names. Some examples of queries include service:<service-name> and <slo-name>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/data-sources/service_level_objectives#query DataDatadogServiceLevelObjectives#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
	// Filter results based on a single SLO tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.0/docs/data-sources/service_level_objectives#tags_query DataDatadogServiceLevelObjectives#tags_query}
	TagsQuery *string `field:"optional" json:"tagsQuery" yaml:"tagsQuery"`
}

