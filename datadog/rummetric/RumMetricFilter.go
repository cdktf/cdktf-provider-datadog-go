// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rummetric


type RumMetricFilter struct {
	// The search query. Follows RUM search syntax.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.63.0/docs/resources/rum_metric#query RumMetric#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

