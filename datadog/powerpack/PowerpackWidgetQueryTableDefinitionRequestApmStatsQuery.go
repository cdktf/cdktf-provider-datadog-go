// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package powerpack


type PowerpackWidgetQueryTableDefinitionRequestApmStatsQuery struct {
	// The environment name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/powerpack#env Powerpack#env}
	Env *string `field:"required" json:"env" yaml:"env"`
	// The operation name associated with the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/powerpack#name Powerpack#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization's host group name and value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/powerpack#primary_tag Powerpack#primary_tag}
	PrimaryTag *string `field:"required" json:"primaryTag" yaml:"primaryTag"`
	// The level of detail for the request. Valid values are `service`, `resource`, `span`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/powerpack#row_type Powerpack#row_type}
	RowType *string `field:"required" json:"rowType" yaml:"rowType"`
	// The service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/powerpack#service Powerpack#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/powerpack#columns Powerpack#columns}
	Columns interface{} `field:"optional" json:"columns" yaml:"columns"`
	// The resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.55.0/docs/resources/powerpack#resource Powerpack#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

