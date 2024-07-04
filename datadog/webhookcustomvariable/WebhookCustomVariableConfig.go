// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webhookcustomvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebhookCustomVariableConfig struct {
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
	// Whether the custom variable is secret or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/webhook_custom_variable#is_secret WebhookCustomVariable#is_secret}
	IsSecret interface{} `field:"required" json:"isSecret" yaml:"isSecret"`
	// The name of the variable. It corresponds with `<CUSTOM_VARIABLE_NAME>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/webhook_custom_variable#name WebhookCustomVariable#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the custom variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/webhook_custom_variable#value WebhookCustomVariable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.40.0/docs/resources/webhook_custom_variable#id WebhookCustomVariable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

