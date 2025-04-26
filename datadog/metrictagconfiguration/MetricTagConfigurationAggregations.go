// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metrictagconfiguration


type MetricTagConfigurationAggregations struct {
	// A space aggregation for use in query. Valid values are `avg`, `max`, `min`, `sum`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/metric_tag_configuration#space MetricTagConfiguration#space}
	Space *string `field:"required" json:"space" yaml:"space"`
	// A time aggregation for use in query. Valid values are `avg`, `count`, `max`, `min`, `sum`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/metric_tag_configuration#time MetricTagConfiguration#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

