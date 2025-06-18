// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package complianceresourceevaluationfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComplianceResourceEvaluationFilterConfig struct {
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
	// The cloud provider of the filter's targeted resource. Only `aws`, `gcp`, or `azure` are considered valid cloud providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/compliance_resource_evaluation_filter#cloud_provider ComplianceResourceEvaluationFilter#cloud_provider}
	CloudProvider *string `field:"required" json:"cloudProvider" yaml:"cloudProvider"`
	// The ID of the of the filter's targeted resource.
	//
	// Different cloud providers target different resource IDs:
	//   - `aws`: account ID
	//   - `gcp`: project ID
	//   - `azure`: subscription ID
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/compliance_resource_evaluation_filter#resource_id ComplianceResourceEvaluationFilter#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// List of tags to filter misconfiguration detections. Each entry should follow the format: "key":"value".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/compliance_resource_evaluation_filter#tags ComplianceResourceEvaluationFilter#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
}

