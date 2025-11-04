// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package incidenttype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IncidentTypeConfig struct {
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
	// Name of the incident type. Must be between 1 and 50 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/incident_type#name IncidentType#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the incident type. The description can have a maximum of 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/incident_type#description IncidentType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether this incident type is the default type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.79.0/docs/resources/incident_type#is_default IncidentType#is_default}
	IsDefault interface{} `field:"optional" json:"isDefault" yaml:"isDefault"`
}

