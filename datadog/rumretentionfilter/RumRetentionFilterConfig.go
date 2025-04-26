// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rumretentionfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RumRetentionFilterConfig struct {
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
	// RUM application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/rum_retention_filter#application_id RumRetentionFilter#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The type of RUM events to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/rum_retention_filter#event_type RumRetentionFilter#event_type}
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The name of a RUM retention filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/rum_retention_filter#name RumRetentionFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The sample rate for a RUM retention filter, between 0 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/rum_retention_filter#sample_rate RumRetentionFilter#sample_rate}
	SampleRate *float64 `field:"required" json:"sampleRate" yaml:"sampleRate"`
	// Whether the retention filter is to be enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/rum_retention_filter#enabled RumRetentionFilter#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The Query string for a RUM retention filter. Defaults to `""`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/rum_retention_filter#query RumRetentionFilter#query}
	Query *string `field:"optional" json:"query" yaml:"query"`
}

