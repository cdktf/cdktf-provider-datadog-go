// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package metrictagconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MetricTagConfigurationConfig struct {
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
	// The metric name for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/metric_tag_configuration#metric_name MetricTagConfiguration#metric_name}
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// The metric's type. This field can't be updated after creation. Valid values are `gauge`, `count`, `rate`, `distribution`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/metric_tag_configuration#metric_type MetricTagConfiguration#metric_type}
	MetricType *string `field:"required" json:"metricType" yaml:"metricType"`
	// A list of tag keys that will be queryable for your metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/metric_tag_configuration#tags MetricTagConfiguration#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
	// aggregations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/metric_tag_configuration#aggregations MetricTagConfiguration#aggregations}
	Aggregations interface{} `field:"optional" json:"aggregations" yaml:"aggregations"`
	// Toggle to include/exclude tags as queryable for your metric.
	//
	// Can only be applied to metrics that have one or more tags configured. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/metric_tag_configuration#exclude_tags_mode MetricTagConfiguration#exclude_tags_mode}
	ExcludeTagsMode interface{} `field:"optional" json:"excludeTagsMode" yaml:"excludeTagsMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/metric_tag_configuration#id MetricTagConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Toggle to include/exclude percentiles for a distribution metric.
	//
	// Defaults to false. Can only be applied to metrics that have a `metric_type` of distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/metric_tag_configuration#include_percentiles MetricTagConfiguration#include_percentiles}
	IncludePercentiles interface{} `field:"optional" json:"includePercentiles" yaml:"includePercentiles"`
}

