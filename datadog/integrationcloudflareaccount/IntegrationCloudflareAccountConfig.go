// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationcloudflareaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationCloudflareAccountConfig struct {
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
	// The API key (or token) for the Cloudflare account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/integration_cloudflare_account#api_key IntegrationCloudflareAccount#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// The name of the Cloudflare account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/integration_cloudflare_account#name IntegrationCloudflareAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The email associated with the Cloudflare account.
	//
	// If an API key is provided (and not a token), this field is also required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/integration_cloudflare_account#email IntegrationCloudflareAccount#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// An allowlist of resources to pull metrics for. Includes `web`, `dns`, `lb` (load balancer), and `worker`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/integration_cloudflare_account#resources IntegrationCloudflareAccount#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

