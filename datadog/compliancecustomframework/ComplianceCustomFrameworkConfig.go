// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package compliancecustomframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComplianceCustomFrameworkConfig struct {
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
	// The framework handle. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/compliance_custom_framework#handle ComplianceCustomFramework#handle}
	Handle *string `field:"required" json:"handle" yaml:"handle"`
	// The framework name. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/compliance_custom_framework#name ComplianceCustomFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The framework version. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/compliance_custom_framework#version ComplianceCustomFramework#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The URL of the icon representing the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/compliance_custom_framework#icon_url ComplianceCustomFramework#icon_url}
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.74.0/docs/resources/compliance_custom_framework#requirements ComplianceCustomFramework#requirements}
	Requirements interface{} `field:"optional" json:"requirements" yaml:"requirements"`
}

