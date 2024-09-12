// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spansmetric


type SpansMetricGroupBy struct {
	// The path to the value the span-based metric will be aggregated over.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/spans_metric#path SpansMetric#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// Eventual name of the tag that gets created. By default, the path attribute is used as the tag name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.44.1/docs/resources/spans_metric#tag_name SpansMetric#tag_name}
	TagName *string `field:"optional" json:"tagName" yaml:"tagName"`
}

