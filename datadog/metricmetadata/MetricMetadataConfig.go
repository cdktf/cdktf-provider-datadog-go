// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metricmetadata

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetricMetadataConfig struct {
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
	// The name of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#metric MetricMetadata#metric}
	Metric *string `field:"required" json:"metric" yaml:"metric"`
	// A description of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#description MetricMetadata#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#id MetricMetadata#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Per unit of the metric such as `second` in `bytes per second`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#per_unit MetricMetadata#per_unit}
	PerUnit *string `field:"optional" json:"perUnit" yaml:"perUnit"`
	// A short name of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#short_name MetricMetadata#short_name}
	ShortName *string `field:"optional" json:"shortName" yaml:"shortName"`
	// If applicable, statsd flush interval in seconds for the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#statsd_interval MetricMetadata#statsd_interval}
	StatsdInterval *float64 `field:"optional" json:"statsdInterval" yaml:"statsdInterval"`
	// Metric type such as `count`, `gauge`, or `rate`.
	//
	// Updating a metric of type `distribution` is not supported. If you would like to see the `distribution` type returned, contact [Datadog support](https://docs.datadoghq.com/help/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#type MetricMetadata#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Primary unit of the metric such as `byte` or `operation`.
	//
	// For a list of allowed units, refer to [Datadog metric unit documentation](https://docs.datadoghq.com/metrics/units/#unit-list).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.80.0/docs/resources/metric_metadata#unit MetricMetadata#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

