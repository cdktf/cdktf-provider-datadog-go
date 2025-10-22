// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workflowautomation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkflowAutomationConfig struct {
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
	// Description of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/workflow_automation#description WorkflowAutomation#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// Name of the workflow. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/workflow_automation#name WorkflowAutomation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Set the workflow to published or unpublished.
	//
	// Workflows in an unpublished state are only executable through manual runs. Automatic triggers such as Schedule do not execute the workflow until it is published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/workflow_automation#published WorkflowAutomation#published}
	Published interface{} `field:"required" json:"published" yaml:"published"`
	// The spec defines what the workflow does.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/workflow_automation#spec_json WorkflowAutomation#spec_json}
	SpecJson *string `field:"required" json:"specJson" yaml:"specJson"`
	// Tags of the workflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/workflow_automation#tags WorkflowAutomation#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
	// If a webhook trigger is defined on this workflow, a webhookSecret is required and should be provided here.
	//
	// String length must be at least 16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.77.0/docs/resources/workflow_automation#webhook_secret WorkflowAutomation#webhook_secret}
	WebhookSecret *string `field:"optional" json:"webhookSecret" yaml:"webhookSecret"`
}

