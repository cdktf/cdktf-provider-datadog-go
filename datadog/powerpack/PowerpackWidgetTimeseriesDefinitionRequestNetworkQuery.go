// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetTimeseriesDefinitionRequestNetworkQuery struct {
	// The name of the index to query.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#index Powerpack#index}
	Index *string `field:"required" json:"index" yaml:"index"`
	// compute_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#compute_query Powerpack#compute_query}
	ComputeQuery *PowerpackWidgetTimeseriesDefinitionRequestNetworkQueryComputeQuery `field:"optional" json:"computeQuery" yaml:"computeQuery"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#group_by Powerpack#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// multi_compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#multi_compute Powerpack#multi_compute}
	MultiCompute interface{} `field:"optional" json:"multiCompute" yaml:"multiCompute"`
	// The search query to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.39.0/docs/resources/powerpack#search_query Powerpack#search_query}
	SearchQuery *string `field:"optional" json:"searchQuery" yaml:"searchQuery"`
}

