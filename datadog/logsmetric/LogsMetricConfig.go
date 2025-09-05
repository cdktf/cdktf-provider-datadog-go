// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logsmetric

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogsMetricConfig struct {
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
	// compute block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_metric#compute LogsMetric#compute}
	Compute *LogsMetricCompute `field:"required" json:"compute" yaml:"compute"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_metric#filter LogsMetric#filter}
	Filter *LogsMetricFilter `field:"required" json:"filter" yaml:"filter"`
	// The name of the log-based metric. This field can't be updated after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_metric#name LogsMetric#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// group_by block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_metric#group_by LogsMetric#group_by}
	GroupBy interface{} `field:"optional" json:"groupBy" yaml:"groupBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.73.0/docs/resources/logs_metric#id LogsMetric#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

