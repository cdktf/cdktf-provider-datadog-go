// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rummetric


type RumMetricGroupBy struct {
	// The path to the value the RUM-based metric will be aggregated over.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/rum_metric#path RumMetric#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Name of the tag that gets created. By default, `path` is used as the tag name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.59.1/docs/resources/rum_metric#tag_name RumMetric#tag_name}
	TagName *string `field:"optional" json:"tagName" yaml:"tagName"`
}

