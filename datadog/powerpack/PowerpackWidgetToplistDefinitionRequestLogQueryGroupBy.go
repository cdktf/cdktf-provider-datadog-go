// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetToplistDefinitionRequestLogQueryGroupBy struct {
	// The facet name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#facet Powerpack#facet}
	Facet *string `field:"optional" json:"facet" yaml:"facet"`
	// The maximum number of items in the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#limit Powerpack#limit}
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// sort_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.67.0/docs/resources/powerpack#sort_query Powerpack#sort_query}
	SortQuery *PowerpackWidgetToplistDefinitionRequestLogQueryGroupBySortQuery `field:"optional" json:"sortQuery" yaml:"sortQuery"`
}

