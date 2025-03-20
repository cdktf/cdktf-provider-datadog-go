// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rummetric


type RumMetricUniqueness struct {
	// When to count updatable events. `match` when the event is first seen, or `end` when the event is complete.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.58.0/docs/resources/rum_metric#when RumMetric#when}
	When *string `field:"optional" json:"when" yaml:"when"`
}

