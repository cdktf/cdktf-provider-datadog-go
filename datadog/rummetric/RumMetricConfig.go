// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rummetric

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RumMetricConfig struct {
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
	// The type of RUM events to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/rum_metric#event_type RumMetric#event_type}
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The name of the RUM-based metric. This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/rum_metric#name RumMetric#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/rum_metric#compute RumMetric#compute}
	Compute *RumMetricCompute `field:"optional" json:"compute" yaml:"compute"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/rum_metric#filter RumMetric#filter}
	Filter *RumMetricFilter `field:"optional" json:"filter" yaml:"filter"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/rum_metric#group_by RumMetric#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// uniqueness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/rum_metric#uniqueness RumMetric#uniqueness}
	Uniqueness *RumMetricUniqueness `field:"optional" json:"uniqueness" yaml:"uniqueness"`
}

