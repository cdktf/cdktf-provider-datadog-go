// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationawseventbridge

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationAwsEventBridgeConfig struct {
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
	// Your AWS Account ID without dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_aws_event_bridge#account_id IntegrationAwsEventBridge#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The given part of the event source name, which is then combined with an assigned suffix to form the full name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_aws_event_bridge#event_generator_name IntegrationAwsEventBridge#event_generator_name}
	EventGeneratorName *string `field:"required" json:"eventGeneratorName" yaml:"eventGeneratorName"`
	// The event source's [AWS region](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_aws_event_bridge#region IntegrationAwsEventBridge#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// True if Datadog should create the event bus in addition to the event source.
	//
	// Requires the `events:CreateEventBus` permission. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.66.0/docs/resources/integration_aws_event_bridge#create_event_bus IntegrationAwsEventBridge#create_event_bus}
	CreateEventBus interface{} `field:"optional" json:"createEventBus" yaml:"createEventBus"`
}

