// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apikey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiKeyConfig struct {
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
	// Name for API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/api_key#name ApiKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether the API key is used for remote config.
	//
	// Set to true only if remote config is enabled in `/organization-settings/remote-config`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.69.0/docs/resources/api_key#remote_config_read_enabled ApiKey#remote_config_read_enabled}
	RemoteConfigReadEnabled interface{} `field:"optional" json:"remoteConfigReadEnabled" yaml:"remoteConfigReadEnabled"`
}

