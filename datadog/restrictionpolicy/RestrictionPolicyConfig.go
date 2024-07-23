// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package restrictionpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RestrictionPolicyConfig struct {
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
	// Identifier for the resource, formatted as resource_type:resource_id.
	//
	// Note: Dashboards support is in private beta. Reach out to your Datadog contact or support to enable this.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/restriction_policy#resource_id RestrictionPolicy#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// bindings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.42.0/docs/resources/restriction_policy#bindings RestrictionPolicy#bindings}
	Bindings interface{} `field:"optional" json:"bindings" yaml:"bindings"`
}

