// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package compliancecustomframework


type ComplianceCustomFrameworkRequirements struct {
	// The name of the requirement. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/compliance_custom_framework#name ComplianceCustomFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.71.0/docs/resources/compliance_custom_framework#controls ComplianceCustomFramework#controls}
	Controls interface{} `field:"optional" json:"controls" yaml:"controls"`
}

