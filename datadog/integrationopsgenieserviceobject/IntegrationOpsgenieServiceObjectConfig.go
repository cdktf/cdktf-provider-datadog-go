// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationopsgenieserviceobject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationOpsgenieServiceObjectConfig struct {
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
	// The name for the Opsgenie service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_opsgenie_service_object#name IntegrationOpsgenieServiceObject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Opsgenie API key for the Opsgenie service.
	//
	// Note: Since the Datadog API never returns Opsgenie API keys, it is impossible to detect [drifts](https://www.hashicorp.com/blog/detecting-and-managing-drift-with-terraform). The best way to solve a drift is to manually mark the Service Object resource with [terraform taint](https://www.terraform.io/docs/commands/taint.html) to have it destroyed and recreated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_opsgenie_service_object#opsgenie_api_key IntegrationOpsgenieServiceObject#opsgenie_api_key}
	OpsgenieApiKey *string `field:"required" json:"opsgenieApiKey" yaml:"opsgenieApiKey"`
	// The region for the Opsgenie service. Valid values are `us`, `eu`, `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_opsgenie_service_object#region IntegrationOpsgenieServiceObject#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// The custom url for a custom region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_opsgenie_service_object#custom_url IntegrationOpsgenieServiceObject#custom_url}
	CustomUrl *string `field:"optional" json:"customUrl" yaml:"customUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.60.1/docs/resources/integration_opsgenie_service_object#id IntegrationOpsgenieServiceObject#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

