// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetHostmapDefinitionRequestSizeProcessQuery struct {
	// Your chosen metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/powerpack#metric Powerpack#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// A list of processes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/powerpack#filter_by Powerpack#filter_by}
	FilterBy *[]*string `field:"optional" json:"filterBy" yaml:"filterBy"`
	// The max number of items in the filter list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/powerpack#limit Powerpack#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Your chosen search term.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.54.0/docs/resources/powerpack#search_by Powerpack#search_by}
	SearchBy *string `field:"optional" json:"searchBy" yaml:"searchBy"`
}

