// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigProcessorsQuota struct {
	// Whether to drop events exceeding the limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#drop_events ObservabilityPipeline#drop_events}
	DropEvents interface{} `field:"required" json:"dropEvents" yaml:"dropEvents"`
	// The unique ID of the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#include ObservabilityPipeline#include}
	Include *string `field:"required" json:"include" yaml:"include"`
	// The inputs for the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// limit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#limit ObservabilityPipeline#limit}
	Limit *ObservabilityPipelineConfigProcessorsQuotaLimit `field:"required" json:"limit" yaml:"limit"`
	// The name of the quota.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#name ObservabilityPipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether to ignore when partition fields are missing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#ignore_when_missing_partitions ObservabilityPipeline#ignore_when_missing_partitions}
	IgnoreWhenMissingPartitions interface{} `field:"optional" json:"ignoreWhenMissingPartitions" yaml:"ignoreWhenMissingPartitions"`
	// The action to take when the quota is exceeded: `drop`, `no_action`, or `overflow_routing`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#overflow_action ObservabilityPipeline#overflow_action}
	OverflowAction *string `field:"optional" json:"overflowAction" yaml:"overflowAction"`
	// overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#overrides ObservabilityPipeline#overrides}
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
	// List of partition fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/observability_pipeline#partition_fields ObservabilityPipeline#partition_fields}
	PartitionFields *[]*string `field:"optional" json:"partitionFields" yaml:"partitionFields"`
}

