// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package compliancecustomframework


type ComplianceCustomFrameworkRequirementsControls struct {
	// The name of the control. String length must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/compliance_custom_framework#name ComplianceCustomFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The set of rules IDs for the control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.70.0/docs/resources/compliance_custom_framework#rules_id ComplianceCustomFramework#rules_id}
	RulesId *[]*string `field:"required" json:"rulesId" yaml:"rulesId"`
}

