// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebhookConfig struct {
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
	// The name of the webhook. It corresponds with `<WEBHOOK_NAME>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/webhook#name Webhook#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL of the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/webhook#url Webhook#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The headers attached to the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/webhook#custom_headers Webhook#custom_headers}
	CustomHeaders *string `field:"optional" json:"customHeaders" yaml:"customHeaders"`
	// Encoding type. Valid values are `json`, `form`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/webhook#encode_as Webhook#encode_as}
	EncodeAs *string `field:"optional" json:"encodeAs" yaml:"encodeAs"`
	// The payload of the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.52.1/docs/resources/webhook#payload Webhook#payload}
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
}

